// Package plugindemo a demo plugin.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	var config Config
	err := json.Unmarshal(handler.Host.GetConfig(), &config)
	if err != nil {
		handler.Host.Log(api.LogLevelError, fmt.Sprintf("Could not load config %v", err))
		os.Exit(1)
	}

	mw, err := New(config)
	if err != nil {
		handler.Host.Log(api.LogLevelError, fmt.Sprintf("Could not load config %v", err))
		os.Exit(1)
	}
	handler.HandleRequestFn = mw.handleRequest
}

// Config the plugin configuration.
type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

// Demo a Demo plugin.
type Demo struct {
	headers map[string]string
}

// New created a new Demo plugin.
func New(config Config) (*Demo, error) {
	if len(config.Headers) == 0 {
		return nil, fmt.Errorf("headers cannot be empty")
	}

	return &Demo{
		headers: config.Headers,
	}, nil
}

func (a *Demo) handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	for key, value := range a.headers {
		req.Headers().Set(key, value)
	}

	return true, 0
}
