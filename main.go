package main

import (
	"net/url"

	pdfparser "github.com/z3f1r/network-go/utils/pdf"
	"github.com/z3f1r/network-go/webcrawler"
)

func main() {
	// examples.TestGetRequest()

	target, _ := url.ParseRequestURI("https://example.com")
	page := webcrawler.GetHTMLPage(*target)

	page.Beatify()

	// fmt.Printf("%s", page.SourceCode)

	pdfPath := "/Users/o.petrakov/Work/pentest/projects/unicredit_web_08_2020/files/dok.pdf"
	pdfparser.TestParsePDF(pdfPath)
}
