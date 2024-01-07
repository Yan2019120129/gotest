package websocket_test_test

import (
	"gotest/middleware/websocket_test/two/server"
	"net/http"
	"testing"
)

// TestHandleFunc 启动websocket
func TestHandleFunc(t *testing.T) {
	http.HandleFunc("/websocket", server.HandleWebSocket)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}
