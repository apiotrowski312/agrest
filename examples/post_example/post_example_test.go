package main

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	agrest "github.com/apiotrowski312/agrest/src"
)

func Test_postReq(t *testing.T) {

	mockClient := agrest.StartMocking()

	mockClient.MockRequest(&http.Request{
		Response: &http.Response{},
		Method:   http.MethodGet, // <-- it should be MethodPost (cause we are doing post request)
		URL: &url.URL{
			Path: "/",
		},
	})

	tests := []struct {
		name    string
		want    *http.Response
		wantErr bool
	}{
		{
			"Basic Post Mock",
			&http.Response{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := postReq()
			if (err != nil) != tt.wantErr {
				t.Errorf("getReq() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getReq() = %v, want %v", got, tt.want)
			}
		})
	}
}
