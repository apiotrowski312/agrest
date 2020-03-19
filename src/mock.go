package agrest

import (
	"errors"
	"fmt"
	"net/http"
)

type mockAgrest struct {
	request map[string]map[string]*http.Request // [method][path]
	path    map[string]*http.Request            // [path]
	method  map[string]*http.Request            // [method]
	all     *http.Request
}

func StartMocking() *mockAgrest {
	mockClient := mockAgrest{}

	client = &mockClient
	return &mockClient
}

func (m mockAgrest) Do(req *http.Request) (*http.Response, error) {

	if req.URL != nil && m.request[req.Method][req.URL.Path] != nil {
		return m.request[req.Method][req.URL.Path].Response, nil
	}

	if req.URL != nil && m.path[req.URL.Path] != nil {
		return m.path[req.URL.Path].Response, nil
	}

	if m.method[req.Method] != nil {
		return m.method[req.Method].Response, nil
	}

	if m.all != nil {
		return m.all.Response, nil
	}

	fmt.Println(m.all)

	return nil, errors.New("No mock set for this path")
}

func (m *mockAgrest) Reset(req *http.Request) {
	m.all = nil
	m.method = nil
	m.request = nil
	m.path = nil
}

func (m *mockAgrest) MockAll(req *http.Request) {
	m.all = req
}

func (m *mockAgrest) MockMethod(req *http.Request) {
	m.method[req.Method] = req
}

func (m *mockAgrest) MockPath(req *http.Request) {
	m.path[req.URL.Path] = req
}

func (m *mockAgrest) MockRequest(req *http.Request) {
	m.request[req.Method][req.URL.Path] = req
}
