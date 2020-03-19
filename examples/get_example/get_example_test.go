package main

import (
	"net/http"
	"reflect"
	"testing"

	agrest "github.com/apiotrowski312/agrest/src"
)

func Test_getReq(t *testing.T) {

	mockClient := agrest.StartMocking()

	mockClient.MockAll(&http.Request{Response: &http.Response{}})

	tests := []struct {
		name    string
		want    *http.Response
		wantErr bool
	}{
		{
			"Basic Get Mock",
			&http.Response{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getReq()
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
