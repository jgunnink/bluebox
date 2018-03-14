package server

import (
	"net/http"
	"testing"

	"github.com/blockninja/ninjarouter"
	"github.com/jgunnink/bluebox"
)

func TestAuthController_SignIn(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
		AuthService bluebox.AuthService
		UserService bluebox.UserService
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
			ac := &AuthController{
				Mux:         tt.fields.Mux,
				AuthService: tt.fields.AuthService,
				UserService: tt.fields.UserService,
			}
			ac.SignIn(tt.args.w, tt.args.r)
		})
	}
}

func TestAuthController_SignOut(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
		AuthService bluebox.AuthService
		UserService bluebox.UserService
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
			ac := &AuthController{
				Mux:         tt.fields.Mux,
				AuthService: tt.fields.AuthService,
				UserService: tt.fields.UserService,
			}
			ac.SignOut(tt.args.w, tt.args.r)
		})
	}
}

func TestAuthController_Check(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
		AuthService bluebox.AuthService
		UserService bluebox.UserService
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
			ac := &AuthController{
				Mux:         tt.fields.Mux,
				AuthService: tt.fields.AuthService,
				UserService: tt.fields.UserService,
			}
			ac.Check(tt.args.w, tt.args.r)
		})
	}
}
