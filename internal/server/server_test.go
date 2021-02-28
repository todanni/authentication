package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/mux"
	"github.com/todanni/template-repository/pkg/example"
)

func TestNewExampleService(t *testing.T) {
	type args struct {
		repo   example.Repository
		router *mux.Router
	}
	tests := []struct {
		name string
		args args
		want example.Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewExampleService(tt.args.repo, tt.args.router); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewExampleService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_exampleService_ExampleMethod(t *testing.T) {
	type fields struct {
		repo   example.Repository
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &exampleService{
				repo:   tt.fields.repo,
				router: tt.fields.router,
			}
		})
	}
}
