package main

import (
	"context"
	"crypto/subtle"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/tus/tusd/v2/pkg/filelocker"
	"github.com/tus/tusd/v2/pkg/filestore"
	tusd "github.com/tus/tusd/v2/pkg/handler"
)

type contextKey string

const subjectKey contextKey = "subject"

func main() {
	var (
		addr          = flag.String("addr", ":8080", "HTTP listen address")
		uploadDir     = flag.String("upload-dir", "./data/uploads", "tusd local storage directory")
		token         = flag.String("token", "demo-token", "demo bearer token")
		maxSize       = flag.Int64("max-size", 2<<30, "maximum upload size in bytes")
		allowedOrigin = flag.String("allowed-origin", "http://localhost:3000", "allowed browser Origin; use * only for local testing")
		allowedExts   = flag.String("allowed-ext", ".zip,.pdf,.png,.jpg,.jpeg,.txt,.csv,.mp4", "comma-separated allowed file extensions")
	)
	flag.Parse()

	if err := os.MkdirAll(*uploadDir, 0o775); err != nil {
		log.Fatalf("create upload directory: %v", err)
	}

	store := filestore.New(*uploadDir)
	locker := filelocker.New(*uploadDir)
	composer := tusd.NewStoreComposer()
	store.UseIn(composer)
	locker.UseIn(composer)

	cors := tusd.DefaultCorsConfig
	cors.AllowOrigin = compileOrigin(*allowedOrigin)
	cors.AllowCredentials = false

	exts := parseAllowedExtensions(*allowedExts)
	handler, err := tusd.NewHandler(tusd.Config{
		BasePath:                "/files/",
		StoreComposer:           composer,
		MaxSize:                 *maxSize,
		DisableDownload:         false,
		DisableTermination:      false,
		DisableConcatenation:    false,
		NotifyCreatedUploads:    true,
		NotifyUploadProgress:    true,
		NotifyCompleteUploads:   true,
		NotifyTerminatedUploads: true,
		UploadProgressInterval:  time.Second,
		RespectForwardedHeaders: true,
		Cors:                    &cors,
		PreUploadCreateCallback: validateUpload(exts),
		PreFinishResponseCallback: func(hook tusd.HookEvent) (tusd.HTTPResponse, error) {
			return tusd.HTTPResponse{Header: tusd.HTTPHeader{
				"X-Upload-Complete": "true",
			}}, nil
		},
		PreUploadTerminateCallback: authorizeTermination,
	})
	if err != nil {
		log.Fatalf("create tusd handler: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	observeEvents(ctx, handler)

	tusHTTPHandler := authMiddleware(*token, http.StripPrefix("/files", handler))
	mux := http.NewServeMux()
	mux.Handle("/files", tusHTTPHandler)
	mux.Handle("/files/", tusHTTPHandler)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	server := &http.Server{
		Addr:              *addr,
		Handler:           requestLog(mux),
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       60 * time.Second,
		// Do not set a short ReadTimeout/WriteTimeout for large uploads.
		// tusd v2 manages active network deadlines through net/http.ResponseController.
	}

	go func() {
		log.Printf("tusd server listening on %s; endpoint=http://localhost%s/files/", *addr, *addr)
		log.Printf("upload directory=%s max-size=%d", *uploadDir, *maxSize)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("serve: %v", err)
		}
	}()

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
}

func validateUpload(allowed map[string]struct{}) func(tusd.HookEvent) (tusd.HTTPResponse, tusd.FileInfoChanges, error) {
	return func(hook tusd.HookEvent) (tusd.HTTPResponse, tusd.FileInfoChanges, error) {
		filename := strings.TrimSpace(hook.Upload.MetaData["filename"])
		if filename == "" {
			return tusd.HTTPResponse{}, tusd.FileInfoChanges{}, tusd.NewError(
				"ERR_FILENAME_REQUIRED",
				"metadata field filename is required",
				http.StatusBadRequest,
			)
		}

		// filepath.Base on Linux does not treat backslashes as separators.
		filename = filepath.Base(strings.ReplaceAll(filename, "\\", "/"))
		ext := strings.ToLower(filepath.Ext(filename))
		if _, ok := allowed[ext]; !ok {
			return tusd.HTTPResponse{}, tusd.FileInfoChanges{}, tusd.NewError(
				"ERR_FILE_TYPE_NOT_ALLOWED",
				fmt.Sprintf("file extension %q is not allowed", ext),
				http.StatusUnsupportedMediaType,
			)
		}

		subject, _ := hook.Context.Value(subjectKey).(string)
		metadata := tusd.MetaData{
			"filename":   filename,
			"owner":      subject,
			"created_at": time.Now().UTC().Format(time.RFC3339Nano),
		}
		if value := strings.TrimSpace(hook.Upload.MetaData["filetype"]); value != "" {
			metadata["filetype"] = value
		}
		if value := strings.TrimSpace(hook.Upload.MetaData["description"]); value != "" {
			metadata["description"] = value
		}

		return tusd.HTTPResponse{Header: tusd.HTTPHeader{
			"X-Upload-Validated": "true",
		}}, tusd.FileInfoChanges{MetaData: metadata}, nil
	}
}

func authorizeTermination(hook tusd.HookEvent) (tusd.HTTPResponse, error) {
	subject, _ := hook.Context.Value(subjectKey).(string)
	owner := hook.Upload.MetaData["owner"]
	if owner == "" || owner != subject {
		return tusd.HTTPResponse{}, tusd.NewError(
			"ERR_UPLOAD_FORBIDDEN",
			"the current user cannot terminate this upload",
			http.StatusForbidden,
		)
	}
	return tusd.HTTPResponse{}, nil
}

func authMiddleware(expectedToken string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Browsers normally send CORS preflight without application credentials.
		if r.Method == http.MethodOptions {
			next.ServeHTTP(w, r)
			return
		}

		const prefix = "Bearer "
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, prefix) || subtle.ConstantTimeCompare(
			[]byte(strings.TrimSpace(strings.TrimPrefix(auth, prefix))),
			[]byte(expectedToken),
		) != 1 {
			w.Header().Set("WWW-Authenticate", "Bearer")
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// TODO(auth): replace the static token with JWT/OIDC validation and use the
		// token's immutable subject claim as the upload owner.
		ctx := context.WithValue(r.Context(), subjectKey, "demo-user")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func observeEvents(ctx context.Context, handler *tusd.Handler) {
	go func() {
		for {
			select {
			case event := <-handler.CreatedUploads:
				log.Printf("created id=%s size=%d filename=%q", event.Upload.ID, event.Upload.Size, event.Upload.MetaData["filename"])
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case event := <-handler.UploadProgress:
				percent := float64(0)
				if event.Upload.Size > 0 {
					percent = float64(event.Upload.Offset) / float64(event.Upload.Size) * 100
				}
				log.Printf("progress id=%s offset=%d/%d %.1f%%", event.Upload.ID, event.Upload.Offset, event.Upload.Size, percent)
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case event := <-handler.CompleteUploads:
				log.Printf("completed id=%s filename=%q path=%q", event.Upload.ID, event.Upload.MetaData["filename"], event.Upload.Storage[filestore.StorageKeyPath])
				// TODO(finalize): enqueue virus scanning/media processing, persist the
				// business file record, and publish the object only after validation.
			case <-ctx.Done():
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case event := <-handler.TerminatedUploads:
				log.Printf("terminated id=%s filename=%q", event.Upload.ID, event.Upload.MetaData["filename"])
			case <-ctx.Done():
				return
			}
		}
	}()
}

func parseAllowedExtensions(value string) map[string]struct{} {
	result := make(map[string]struct{})
	for _, item := range strings.Split(value, ",") {
		ext := strings.ToLower(strings.TrimSpace(item))
		if ext == "" {
			continue
		}
		if !strings.HasPrefix(ext, ".") {
			ext = "." + ext
		}
		result[ext] = struct{}{}
	}
	return result
}

func compileOrigin(origin string) *regexp.Regexp {
	if origin == "*" {
		return regexp.MustCompile(".*")
	}
	return regexp.MustCompile("^" + regexp.QuoteMeta(origin) + "$")
}

func requestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		started := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("request method=%s path=%s elapsed=%s", r.Method, r.URL.Path, time.Since(started).Round(time.Millisecond))
	})
}
