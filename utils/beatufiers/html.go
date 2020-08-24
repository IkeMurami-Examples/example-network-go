package beatufiers

import "github.com/yosssi/gohtml"

func HTMLBeatify(code string) string {
	return gohtml.Format(code)
}
