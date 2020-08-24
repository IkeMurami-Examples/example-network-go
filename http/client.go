package http

import (
	"net/http"
	"time"
)

// NewHTTPClient - get http.Client
func NewHTTPClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: transport}

	return client
}
