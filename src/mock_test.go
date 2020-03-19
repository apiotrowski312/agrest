package agrest

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestMockAgrest_Do(t *testing.T) {
	type fields struct {
		request map[string]map[string]*http.Request
		path    map[string]*http.Request
		method  map[string]*http.Request
		all     *http.Request
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *http.Response
		wantErr bool
	}{
		{
			"First test",
			fields{
				all: &http.Request{
					Response: &http.Response{
						StatusCode: 200,
					},
				},
			},
			args{req: &http.Request{}},
			&http.Response{
				StatusCode: 200,
			},
			false,
		},
		{
			"Nil test, should throw error",
			fields{all: nil},
			args{req: &http.Request{}},
			nil,
			true,
		},
		{
			"Method test",
			fields{
				method: map[string]*http.Request{
					http.MethodGet: &http.Request{
						Response: &http.Response{
							StatusCode: 200,
						},
					},
				},
			},
			args{req: &http.Request{
				Method: http.MethodGet,
			}},
			&http.Response{
				StatusCode: 200,
			},
			false,
		},
		{
			"Path test",
			fields{
				path: map[string]*http.Request{
					"/test": &http.Request{
						Response: &http.Response{
							StatusCode: 200,
						},
					},
				},
			},
			args{req: &http.Request{
				URL: &url.URL{
					Path: "/test",
				},
			}},
			&http.Response{
				StatusCode: 200,
			},
			false,
		},
		{
			"Requestr test",
			fields{
				request: map[string]map[string]*http.Request{
					http.MethodGet: map[string]*http.Request{
						"/test": &http.Request{
							Response: &http.Response{
								StatusCode: 200,
							},
						},
					},
				},
			},
			args{req: &http.Request{
				Method: http.MethodGet,
				URL: &url.URL{
					Path: "/test",
				},
			}},
			&http.Response{
				StatusCode: 200,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := mockAgrest{
				request: tt.fields.request,
				path:    tt.fields.path,
				method:  tt.fields.method,
				all:     tt.fields.all,
			}
			got, err := m.Do(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("mockAgrest.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mockAgrest.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
