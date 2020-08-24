package main

import (
	"net/url"

	"./webcrawler"
)

func main() {
	// examples.TestGetRequest()

	target, _ := url.ParseRequestURI("https://example.com")
	webcrawler.GetHTMLPage(*target)
}
