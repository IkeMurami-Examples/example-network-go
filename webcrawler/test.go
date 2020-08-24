package webcrawler

import (
	"fmt"
	"io/ioutil"
	"net/url"

	http_client "../http"
)

// GetHTMLPage - request some html by URL
func GetHTMLPage(target url.URL) {
	client := http_client.NewHTTPClient()
	resp, _ := client.Get(target.String())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", body)
}
