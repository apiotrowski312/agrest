package agrest

import "net/http"

func (a agrest) Get(path string) (*http.Response, error) {
	req, err := a.get(path)

	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

func (a agrest) get(path string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, a.baseURL+path, nil)
	if err != nil {
		return nil, err
	}

	return request, nil
}
