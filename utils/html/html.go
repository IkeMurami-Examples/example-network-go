package html

import (
	"net/url"

	"github.com/yosssi/gohtml"

	"github.com/z3f1r/network-go/utils/js"
)

// Page -
type Page struct {
	SourceCode string
	URL        url.URL
}

// Beatify -
func (page Page) Beatify() Page {
	page.SourceCode = gohtml.Format(page.SourceCode)
	return page
}

// GetJSScripts -
func (page Page) GetJSScripts() []js.Script {
	// use colly for crawling site: https://github.com/gocolly/colly
}
