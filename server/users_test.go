package server

import (
	"net/http"
	"testing"

	"github.com/blockninja/ninjarouter"
	"github.com/jgunnink/bluebox"
)

func TestUserController_UserCreate(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserCreate(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserByID(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserByID(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserAll(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserAll(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserUpdate(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserUpdate(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserArchive(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserArchive(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserUnarchive(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserUnarchive(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserDisable(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserDisable(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserEnable(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserEnable(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserSetResetToken(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserSetResetToken(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_UserPasswordReset(t *testing.T) {
	type fields struct {
		Mux         *ninjarouter.Mux
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
			uc := &UserController{
				Mux:         tt.fields.Mux,
				UserService: tt.fields.UserService,
			}
			uc.UserPasswordReset(tt.args.w, tt.args.r)
		})
	}
}
