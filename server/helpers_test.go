package server

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/gorilla/sessions"
	"github.com/jgunnink/bluebox"
)

func TestLoadCookie(t *testing.T) {
	type args struct {
		r  *http.Request
		cs *sessions.CookieStore
	}
	tests := []struct {
		name    string
		args    args
		want    *bluebox.Cookie
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadCookie(tt.args.r, tt.args.cs)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadCookie() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_marshalAndRespond(t *testing.T) {
	type args struct {
		w      http.ResponseWriter
		result interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marshalAndRespond(tt.args.w, tt.args.result)
		})
	}
}

func Test_mustGetID(t *testing.T) {
	type args struct {
		r    *http.Request
		name string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mustGetID(tt.args.r, tt.args.name); got != tt.want {
				t.Errorf("mustGetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_updateCookieSession(t *testing.T) {
	type args struct {
		cookie  *bluebox.Cookie
		session *sessions.Session
	}
	tests := []struct {
		name string
		args args
		want *sessions.Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateCookieSession(tt.args.cookie, tt.args.session); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateCookieSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
