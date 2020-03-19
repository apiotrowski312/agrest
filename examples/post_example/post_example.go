package main

import (
	"fmt"
	"net/http"

	agrest "github.com/apiotrowski312/agrest/src"
)

func main() {
	req, err := postReq()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(req)
	}
}

func postReq() (*http.Response, error) {
	Client := agrest.CreateClient()
	Client.SetBaseURL("http://google.com")

	req, err := Client.Post("/").AddStringBody("test text", "text/plain").SetHeaders(&http.Header{
		"X-Test": []string{"True"},
	}).Do()

	if err != nil {
		return nil, err
	}

	return req, nil
}
