package http_examples

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// NewClient - get http.client
func NewClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: transport}

	return client
}

// TestGetRequest - get https://example.com/
func TestGetRequest() {
	client := NewClient()

	resp, err := client.Get("https://example.com")
	if err != nil {
		// Handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
