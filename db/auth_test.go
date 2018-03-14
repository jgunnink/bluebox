package db

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx"
	"github.com/jgunnink/bluebox"
)

func TestAuthService_UserSetToken(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	type args struct {
		id           int
		sessionToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := &AuthService{
				db: tt.fields.db,
			}
			got, err := as.UserSetToken(tt.args.id, tt.args.sessionToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthService.UserSetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthService.UserSetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_SignOut(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			as := &AuthService{
				db: tt.fields.db,
			}
			if err := as.SignOut(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("AuthService.SignOut() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
