package db

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx"
)

func TestDB_Migrate(t *testing.T) {
	type fields struct {
		client      *pgx.ConnPool
		authService *AuthService
		userService *UserService
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				client:      tt.fields.client,
				authService: tt.fields.authService,
				userService: tt.fields.userService,
			}
			db.Migrate()
		})
	}
}

func TestDB_Drop(t *testing.T) {
	type fields struct {
		client      *pgx.ConnPool
		authService *AuthService
		userService *UserService
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				client:      tt.fields.client,
				authService: tt.fields.authService,
				userService: tt.fields.userService,
			}
			db.Drop()
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		dbname   string
		username string
		password string
		host     string
		port     uint
	}
	tests := []struct {
		name string
		args args
		want *DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.dbname, tt.args.username, tt.args.password, tt.args.host, tt.args.port); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
