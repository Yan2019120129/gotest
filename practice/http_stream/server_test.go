package http_stream

import (
	"context"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStreamClientReceivesMessages(t *testing.T) {
	server := httptest.NewServer(NewServeMux())
	defer server.Close()

	var received []Message
	client := &Client{HTTPClient: server.Client()}
	err := client.Stream(context.Background(), server.URL+"/stream?count=3&interval=1ms", func(msg Message) error {
		received = append(received, msg)
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(received) != 3 {
		t.Fatalf("received %d messages, want 3", len(received))
	}
	if !received[2].Done {
		t.Fatal("last message should be marked done")
	}
	for i, msg := range received {
		if msg.Index != i+1 {
			t.Fatalf("message index=%d, want %d", msg.Index, i+1)
		}
	}
}

func TestStreamClientCanCancel(t *testing.T) {
	server := httptest.NewServer(NewServeMux())
	defer server.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	received := 0
	client := &Client{HTTPClient: server.Client()}
	err := client.Stream(ctx, server.URL+"/stream?count=10&interval=20ms", func(msg Message) error {
		received++
		cancel()
		return nil
	})
	if err == nil {
		t.Fatal("expected cancellation error")
	}
	if received != 1 {
		t.Fatalf("received %d messages before cancel, want 1", received)
	}

	time.Sleep(5 * time.Millisecond)
}
