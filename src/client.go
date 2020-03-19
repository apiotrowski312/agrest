package agrest

import "net/http"

type agrest struct {
	baseURL string
	timeout int
}

func CreateClient() agrest {
	return agrest{}
}

func (a *agrest) SetBaseURL(baseURL string) {
	a.baseURL = baseURL
}

func (a agrest) Get(path string) *request {
	return createRequest(http.MethodGet, a.baseURL+path)
}

func (a agrest) Post(path string) *request {
	return createRequest(http.MethodPost, a.baseURL+path)
}
