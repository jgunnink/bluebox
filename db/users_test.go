package db

import (
	"reflect"
	"testing"

	"github.com/jackc/pgx"
	"github.com/jgunnink/bluebox"
)

func TestUserService_UserCreate(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	type args struct {
		user *bluebox.User
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
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserCreate(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserCreate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserAll(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[int]*bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserArchive(t *testing.T) {
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
		want    *bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserArchive(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserArchive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserArchive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserUnarchive(t *testing.T) {
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
		want    *bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserUnarchive(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserUnarchive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserUnarchive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserDisable(t *testing.T) {
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
		want    *bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserDisable(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserDisable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserDisable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserEnable(t *testing.T) {
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
		want    *bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserEnable(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserEnable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserEnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserByEmail(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	type args struct {
		email string
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
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserByID(t *testing.T) {
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
		want    *bluebox.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserSetToken(t *testing.T) {
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
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserSetToken(tt.args.id, tt.args.sessionToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserSetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserSetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserUpdate(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	type args struct {
		user *bluebox.User
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
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserUpdate(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserUpdate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserUpdate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserSetResetToken(t *testing.T) {
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
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserSetResetToken(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserSetResetToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("UserService.UserSetResetToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_UserSetPassword(t *testing.T) {
	type fields struct {
		db *pgx.ConnPool
	}
	type args struct {
		id             int
		hashedPassword string
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
			us := &UserService{
				db: tt.fields.db,
			}
			got, err := us.UserSetPassword(tt.args.id, tt.args.hashedPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.UserSetPassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.UserSetPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
