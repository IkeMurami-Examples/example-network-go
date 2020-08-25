package js

import (
	"net/url"

	"github.com/ditashi/jsbeautifier-go/jsbeautifier"
	"github.com/ditashi/jsbeautifier-go/optargs"
)

// Script -
type Script struct {
	SourceCode string
	URL        url.URL
}

// Beatify -
func (code Script) Beatify() Script {
	code.SourceCode = jsbeautifier.Beautify(&code.SourceCode, optargs.MapType{})
	return code
}
