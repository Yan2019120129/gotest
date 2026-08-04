package tusclient

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

const tusVersion = "1.0.0"

type Client struct {
	Endpoint        string
	Token           string
	ChunkSize       int64
	MaxRetries      int
	HTTPClient      *http.Client
	CheckpointStore *CheckpointStore
}

type UploadInfo struct {
	URL      string
	Offset   int64
	Length   int64
	Metadata map[string]string
}

type ServerOptions struct {
	Versions   string
	Extensions string
	MaxSize    string
}

type UploadOptions struct {
	Restart     bool
	Description string
	Progress    func(uploaded, total int64)
}

func New(endpoint, token string, chunkSize int64, maxRetries int, checkpointDir string) (*Client, error) {
	if chunkSize <= 0 {
		return nil, errors.New("chunk size must be positive")
	}
	if maxRetries < 0 {
		return nil, errors.New("max retries cannot be negative")
	}
	endpoint = strings.TrimRight(endpoint, "/") + "/"
	if _, err := url.ParseRequestURI(endpoint); err != nil {
		return nil, fmt.Errorf("invalid endpoint: %w", err)
	}

	return &Client{
		Endpoint:   endpoint,
		Token:      token,
		ChunkSize:  chunkSize,
		MaxRetries: maxRetries,
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				Proxy:                 http.ProxyFromEnvironment,
				MaxIdleConns:          20,
				MaxIdleConnsPerHost:   10,
				IdleConnTimeout:       90 * time.Second,
				ResponseHeaderTimeout: 30 * time.Second,
				ExpectContinueTimeout: time.Second,
			},
			// No global Timeout: a large upload may legitimately take a long time.
		},
		CheckpointStore: NewCheckpointStore(checkpointDir),
	}, nil
}

func (c *Client) Options(ctx context.Context) (ServerOptions, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodOptions, c.Endpoint, nil)
	if err != nil {
		return ServerOptions{}, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return ServerOptions{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		return ServerOptions{}, responseError(resp)
	}
	return ServerOptions{
		Versions:   resp.Header.Get("Tus-Version"),
		Extensions: resp.Header.Get("Tus-Extension"),
		MaxSize:    resp.Header.Get("Tus-Max-Size"),
	}, nil
}

func (c *Client) UploadFile(ctx context.Context, filePath string, options UploadOptions) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("open file: %w", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("stat file: %w", err)
	}
	if !stat.Mode().IsRegular() {
		return "", errors.New("only regular files are supported")
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", fmt.Errorf("absolute file path: %w", err)
	}
	fingerprint := FileFingerprint{
		Path:            absPath,
		Size:            stat.Size(),
		ModTimeUnixNano: stat.ModTime().UnixNano(),
	}

	var checkpoint *Checkpoint
	if !options.Restart {
		checkpoint, err = c.CheckpointStore.Load(fingerprint)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("load checkpoint: %w", err)
		}
	} else {
		_ = c.CheckpointStore.Remove(fingerprint)
	}

	metadata := map[string]string{
		"filename": filepath.Base(filePath),
	}
	if contentType := mime.TypeByExtension(strings.ToLower(filepath.Ext(filePath))); contentType != "" {
		metadata["filetype"] = contentType
	}
	if options.Description != "" {
		metadata["description"] = options.Description
	}

	var uploadURL string
	if checkpoint != nil {
		uploadURL = checkpoint.UploadURL
	} else {
		uploadURL, err = c.createUpload(ctx, stat.Size(), metadata)
		if err != nil {
			return "", err
		}
		checkpoint = &Checkpoint{
			Version:     1,
			Fingerprint: fingerprint,
			UploadURL:   uploadURL,
			Offset:      0,
			Metadata:    metadata,
		}
		if err := c.CheckpointStore.Save(*checkpoint); err != nil {
			return "", fmt.Errorf("save initial checkpoint: %w", err)
		}
	}

	info, err := c.Head(ctx, uploadURL)
	if err != nil {
		return uploadURL, fmt.Errorf("query remote offset: %w", err)
	}
	if info.Length >= 0 && info.Length != stat.Size() {
		return uploadURL, fmt.Errorf("remote length %d differs from local size %d", info.Length, stat.Size())
	}
	if info.Offset > stat.Size() {
		return uploadURL, fmt.Errorf("remote offset %d exceeds local size %d", info.Offset, stat.Size())
	}

	offset := info.Offset
	checkpoint.Offset = offset
	if err := c.CheckpointStore.Save(*checkpoint); err != nil {
		return uploadURL, fmt.Errorf("save synchronized checkpoint: %w", err)
	}
	if options.Progress != nil {
		options.Progress(offset, stat.Size())
	}

	for offset < stat.Size() {
		if err := ctx.Err(); err != nil {
			return uploadURL, err
		}

		chunkLength := min64(c.ChunkSize, stat.Size()-offset)
		newOffset, err := c.patchWithRetry(ctx, file, uploadURL, offset, chunkLength)
		if err != nil {
			return uploadURL, err
		}
		if newOffset <= offset {
			return uploadURL, fmt.Errorf("server did not advance offset: old=%d new=%d", offset, newOffset)
		}

		offset = newOffset
		checkpoint.Offset = offset
		if err := c.CheckpointStore.Save(*checkpoint); err != nil {
			return uploadURL, fmt.Errorf("save checkpoint: %w", err)
		}
		if options.Progress != nil {
			options.Progress(offset, stat.Size())
		}
	}

	if err := c.CheckpointStore.Remove(fingerprint); err != nil && !errors.Is(err, os.ErrNotExist) {
		return uploadURL, fmt.Errorf("remove completed checkpoint: %w", err)
	}
	return uploadURL, nil
}

func (c *Client) Head(ctx context.Context, uploadURL string) (UploadInfo, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodHead, uploadURL, nil)
	if err != nil {
		return UploadInfo{}, err
	}
	c.setTusHeaders(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return UploadInfo{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return UploadInfo{}, responseError(resp)
	}

	offset, err := parseRequiredInt64Header(resp.Header, "Upload-Offset")
	if err != nil {
		return UploadInfo{}, err
	}
	length := int64(-1)
	if raw := resp.Header.Get("Upload-Length"); raw != "" {
		length, err = strconv.ParseInt(raw, 10, 64)
		if err != nil || length < 0 {
			return UploadInfo{}, fmt.Errorf("invalid Upload-Length %q", raw)
		}
	}

	return UploadInfo{
		URL:      uploadURL,
		Offset:   offset,
		Length:   length,
		Metadata: decodeMetadata(resp.Header.Get("Upload-Metadata")),
	}, nil
}

func (c *Client) Cancel(ctx context.Context, uploadURL string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, uploadURL, nil)
	if err != nil {
		return err
	}
	c.setTusHeaders(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		return responseError(resp)
	}
	return nil
}

func (c *Client) Download(ctx context.Context, uploadURL, destination string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uploadURL, nil)
	if err != nil {
		return err
	}
	c.setTusHeaders(req)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return responseError(resp)
	}

	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil && filepath.Dir(destination) != "." {
		return fmt.Errorf("create destination directory: %w", err)
	}
	tmp := destination + ".part"
	file, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("create destination: %w", err)
	}
	_, copyErr := io.Copy(file, resp.Body)
	closeErr := file.Close()
	if copyErr != nil {
		return fmt.Errorf("download body: %w", copyErr)
	}
	if closeErr != nil {
		return fmt.Errorf("close destination: %w", closeErr)
	}
	if err := os.Rename(tmp, destination); err != nil {
		return fmt.Errorf("publish destination: %w", err)
	}
	return nil
}

func (c *Client) createUpload(ctx context.Context, size int64, metadata map[string]string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.Endpoint, bytes.NewReader(nil))
	if err != nil {
		return "", err
	}
	c.setTusHeaders(req)
	req.Header.Set("Upload-Length", strconv.FormatInt(size, 10))
	req.Header.Set("Upload-Metadata", encodeMetadata(metadata))
	req.ContentLength = 0

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("create upload: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return "", responseError(resp)
	}
	location := resp.Header.Get("Location")
	if location == "" {
		return "", errors.New("create response is missing Location header")
	}
	resolved, err := req.URL.Parse(location)
	if err != nil {
		return "", fmt.Errorf("resolve upload URL: %w", err)
	}
	return resolved.String(), nil
}

func (c *Client) patchWithRetry(ctx context.Context, file *os.File, uploadURL string, offset, chunkLength int64) (int64, error) {
	currentOffset := offset
	remaining := chunkLength

	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		newOffset, err := c.patch(ctx, file, uploadURL, currentOffset, remaining)
		if err == nil {
			return newOffset, nil
		}
		if ctx.Err() != nil {
			return currentOffset, ctx.Err()
		}
		if !isRetryable(err) || attempt == c.MaxRetries {
			return currentOffset, fmt.Errorf("patch offset %d after %d attempt(s): %w", currentOffset, attempt+1, err)
		}

		// A broken connection may still have committed bytes on the server. HEAD is
		// authoritative; resume from the server's actual offset rather than replaying.
		info, headErr := c.Head(ctx, uploadURL)
		if headErr == nil {
			if info.Offset < offset || info.Offset > offset+chunkLength {
				return currentOffset, fmt.Errorf("unexpected remote offset after retry: %d", info.Offset)
			}
			currentOffset = info.Offset
			remaining = offset + chunkLength - currentOffset
			if remaining == 0 {
				return currentOffset, nil
			}
		}

		delay := retryDelay(attempt)
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return currentOffset, ctx.Err()
		case <-timer.C:
		}
	}
	return currentOffset, errors.New("unreachable retry state")
}

func (c *Client) patch(ctx context.Context, file *os.File, uploadURL string, offset, length int64) (int64, error) {
	body := io.NewSectionReader(file, offset, length)
	req, err := http.NewRequestWithContext(ctx, http.MethodPatch, uploadURL, body)
	if err != nil {
		return offset, err
	}
	c.setTusHeaders(req)
	req.Header.Set("Content-Type", "application/offset+octet-stream")
	req.Header.Set("Upload-Offset", strconv.FormatInt(offset, 10))
	req.ContentLength = length

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return offset, retryableError{err: err}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusNoContent {
		err := responseError(resp)
		if resp.StatusCode == http.StatusConflict || resp.StatusCode == http.StatusLocked || resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500 {
			return offset, retryableError{err: err}
		}
		return offset, err
	}
	return parseRequiredInt64Header(resp.Header, "Upload-Offset")
}

func (c *Client) setTusHeaders(req *http.Request) {
	req.Header.Set("Tus-Resumable", tusVersion)
	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
}

type retryableError struct{ err error }

func (e retryableError) Error() string { return e.err.Error() }
func (e retryableError) Unwrap() error { return e.err }

func isRetryable(err error) bool {
	var target retryableError
	return errors.As(err, &target)
}

func responseError(resp *http.Response) error {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, 64<<10))
	return fmt.Errorf("HTTP %d %s: %s", resp.StatusCode, resp.Status, strings.TrimSpace(string(body)))
}

func parseRequiredInt64Header(header http.Header, key string) (int64, error) {
	raw := header.Get(key)
	value, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || value < 0 {
		return 0, fmt.Errorf("missing or invalid %s header %q", key, raw)
	}
	return value, nil
}

func encodeMetadata(metadata map[string]string) string {
	keys := make([]string, 0, len(metadata))
	for key := range metadata {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	parts := make([]string, 0, len(keys))
	for _, key := range keys {
		value := base64.StdEncoding.EncodeToString([]byte(metadata[key]))
		parts = append(parts, key+" "+value)
	}
	return strings.Join(parts, ",")
}

func decodeMetadata(raw string) map[string]string {
	result := make(map[string]string)
	for _, item := range strings.Split(raw, ",") {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		parts := strings.SplitN(item, " ", 2)
		if len(parts) == 1 {
			result[parts[0]] = ""
			continue
		}
		decoded, err := base64.StdEncoding.DecodeString(parts[1])
		if err == nil {
			result[parts[0]] = string(decoded)
		}
	}
	return result
}

func retryDelay(attempt int) time.Duration {
	delay := time.Second << minInt(attempt, 3)
	return delay
}

func min64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
