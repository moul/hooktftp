package hooks

import (
	"io"
	"net/http"
	"fmt"
)

var UrlHook = HookComponents{
	func(url string) (io.ReadCloser, error) {

		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		if res.StatusCode != 200 {
			return nil, fmt.Errorf("Bad response '%v' from %v", res.Status, url)
		}

		return res.Body, nil

	},
	func(s string) string {
		return s
	},
}
