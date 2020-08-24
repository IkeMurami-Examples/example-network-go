package beatufiers

import "github.com/ditashi/jsbeautifier-go/jsbeautifier"

func JSBeatify(code string) string {
	code := jsbeautifier.BeautifyFile(file, options)
}
