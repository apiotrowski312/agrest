package agrest

import (
	"io"
	"net/http"
	"strings"
)

type request struct {
	method      string
	url         string
	headers     *http.Header
	body        io.Reader
	contentType [2]string
}

func createRequest(method string, url string) *request {
	return &request{
		method:  method,
		url:     url,
		headers: &http.Header{},
	}
}

func (r *request) SetHeaders(headers *http.Header) *request {
	r.headers = headers

	return r
}

func (r *request) AddStringBody(text string, contentType string) *request {
	r.body = strings.NewReader(text)
	r.contentType = [2]string{"Content-Type", contentType}

	return r
}

func (r *request) AddBody(body io.Reader, contentType string) *request {
	r.body = body
	r.contentType = [2]string{"Content-Type", contentType}

	return r
}

func (r request) Do() (*http.Response, error) {
	req, err := http.NewRequest(r.method, r.url, r.body)
	if err != nil {
		return nil, err
	}

	req.Header = *r.headers

	if r.contentType[0] != "" {
		req.Header.Set(r.contentType[0], r.contentType[1])
	}

	return client.Do(req)
}
