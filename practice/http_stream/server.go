package http_stream

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const (
	defaultStreamCount    = 10
	defaultStreamInterval = 500 * time.Millisecond
)

// StreamHandler writes newline-delimited JSON messages as a chunked HTTP response.
func StreamHandler(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "streaming unsupported", http.StatusInternalServerError)
		return
	}

	count := queryInt(r, "count", defaultStreamCount)
	interval := queryDuration(r, "interval", defaultStreamInterval)

	w.Header().Set("Content-Type", "application/x-ndjson; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	enc := json.NewEncoder(w)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	ctx := r.Context()
	for i := 1; i <= count; i++ {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msg := Message{
			Index: i,
			Text:  fmt.Sprintf("stream message %d", i),
			Time:  time.Now(),
			Done:  i == count,
		}
		if err := enc.Encode(msg); err != nil {
			return
		}
		flusher.Flush()

		if i == count {
			return
		}

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
	}
}

func NewServeMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/stream", StreamHandler)
	mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	return mux
}

func queryInt(r *http.Request, key string, fallback int) int {
	value := r.URL.Query().Get(key)
	if value == "" {
		return fallback
	}
	n, err := strconv.Atoi(value)
	if err != nil || n <= 0 {
		return fallback
	}
	return n
}

func queryDuration(r *http.Request, key string, fallback time.Duration) time.Duration {
	value := r.URL.Query().Get(key)
	if value == "" {
		return fallback
	}
	d, err := time.ParseDuration(value)
	if err == nil && d > 0 {
		return d
	}
	ms, err := strconv.Atoi(value)
	if err != nil || ms <= 0 {
		return fallback
	}
	return time.Duration(ms) * time.Millisecond
}
