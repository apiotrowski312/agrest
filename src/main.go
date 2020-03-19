package agrest

import "net/http"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	client HTTPClient
)

func init() {
	client = &http.Client{}
}
