package agrest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_SetHeaders(t *testing.T) {
	req := createRequest(http.MethodGet, "localhost:8080/")
	req.SetHeaders(&http.Header{
		"test": []string{
			"1", "2", "3",
		},
	})

	reverseHeader := &http.Header{
		"test": []string{
			"3", "2", "1",
		},
	}

	req.SetHeaders(reverseHeader)

	if req.headers != reverseHeader {
		t.Errorf("Smth went wrong with overwriting headers - expected: %v results: %v", reverseHeader, req.headers)
	}
}

func Test_Do(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(261)
		rw.Write([]byte("261 - Test working"))
	}))

	defer server.Close()

	req := createRequest(http.MethodGet, server.URL)

	response, err := req.Do()

	if err != nil {
		t.Errorf("Error with Do request: %v", err)
		return
	}
	if response.StatusCode != 261 {
		t.Errorf("Error with Do request: %v", response)
	}
}
