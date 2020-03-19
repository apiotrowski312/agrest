package main

import (
	"fmt"
	"net/http"

	agrest "github.com/apiotrowski312/agrest/src"
)

func main() {
	req, err := getReq()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(req)
	}
}

func getReq() (*http.Response, error) {
	Client := agrest.CreateClient()
	Client.SetBaseURL("http://google.com")

	req, err := Client.Get("/")

	if err != nil {
		return nil, err
	}

	return req, nil
}
