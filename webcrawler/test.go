package webcrawler

import (
	"io/ioutil"
	"net/url"

	http_client "github.com/z3f1r/network-go/http"
	"github.com/z3f1r/network-go/utils/html"
)

// GetHTMLPage - request some html by URL
func GetHTMLPage(target url.URL) html.Page {

	client := http_client.NewHTTPClient()
	resp, _ := client.Get(target.String())

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return html.Page{URL: target, SourceCode: string(body)}
}
