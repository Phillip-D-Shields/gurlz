package internal

import (
	"time"
)

type Collection struct {
	ID          string    `json:"id" yaml:"id"`
	Name        string    `json:"name" yaml:"name"`
	Description string    `json:"description,omitempty" yaml:"description,omitempty"`
	Requests    []Request `json:"requests,omitempty" yaml:"requests,omitempty"`
	CreatedAt   time.Time `json:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" yaml:"updated_at"`
}

type Config struct {
	DefaultHeaders  map[string]string `json:"default_headers,omitempty" yaml:"default_headers,omitempty"`
	Timeout         string            `json:"timeout" yaml:"timeout"`
	FollowRedirect  bool              `json:"follow_redirect" yaml:"follow_redirect"`
	SaveResponses   bool              `json:"save_responses" yaml:"save_responses"`
	OutputFormat    string            `json:"output_format" yaml:"output_format"`
	ColorOutput     bool              `json:"color_output" yaml:"color_output"`
	DefaultMethod   string            `json:"default_method" yaml:"default_method"`
	MaxResponseSize int64             `json:"max_response_size" yaml:"max_response_size"`
}

type Request struct {
	ID        string            `json:"id" yaml:"id"`
	Name      string            `json:"name" yaml:"name"`
	URL       string            `json:"url" yaml:"url"`
	Method    string            `json:"method" yaml:"method"`
	Headers   map[string]string `json:"headers,omitempty" yaml:"headers,omitempty"`
	Body      string            `json:"body,omitempty" yaml:"body,omitempty"`
	CreatedAt time.Time         `json:"created_at" yaml:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" yaml:"updated_at"`
}

type RequestStore struct {
	Requests []Request `yaml:"requests"`
}

type Response struct {
	StatusCode   int               `json:"status_code"`
	Status       string            `json:"status"`
	Headers      map[string]string `json:"headers"`
	Body         string            `json:"body,omitempty"`
	ResponseTime time.Duration     `json:"response_time"`
	Timestamp    time.Time         `json:"timestamp"`
	RequestID    string            `json:"request_id"`
}
