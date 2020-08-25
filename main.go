package main

import (
	"fmt"
	"net/url"

	"github.com/z3f1r/network-go/webcrawler"
)

func main() {
	// examples.TestGetRequest()

	target, _ := url.ParseRequestURI("https://example.com")
	page := webcrawler.GetHTMLPage(*target)

	page.Beatify()

	fmt.Printf("%s", page.SourceCode)
}
