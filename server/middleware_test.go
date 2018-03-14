package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jgunnink/bluebox"
)

func TestMiddleware_AdminChain(t *testing.T) {
	type fields struct {
		UserService bluebox.UserService
	}
	type args struct {
		h http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(http.ResponseWriter, *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := &Middleware{
				UserService: tt.fields.UserService,
			}
			if got := mw.AdminChain(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middleware.AdminChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleware_InsecureChain(t *testing.T) {
	type fields struct {
		UserService bluebox.UserService
	}
	type args struct {
		h http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(http.ResponseWriter, *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := &Middleware{
				UserService: tt.fields.UserService,
			}
			if got := mw.InsecureChain(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middleware.InsecureChain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleware_SecureChain(t *testing.T) {
	type fields struct {
		UserService bluebox.UserService
	}
	type args struct {
		h http.HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   func(http.ResponseWriter, *http.Request)
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := &Middleware{
				UserService: tt.fields.UserService,
			}
			if got := mw.SecureChain(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Middleware.SecureChain() = %v, want %v", got, tt.want)
			}
		})
	}
}
