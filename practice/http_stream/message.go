package http_stream

import "time"

// Message is one chunk in the HTTP stream.
type Message struct {
	Index int       `json:"index"`
	Text  string    `json:"text"`
	Time  time.Time `json:"time"`
	Done  bool      `json:"done"`
}
