package seed

import (
	"fmt"
	"log"
	"strings"

	"sync"

	"github.com/jgunnink/bluebox"
	"github.com/jgunnink/bluebox/db"
	"github.com/jgunnink/bluebox/helpers"
	"github.com/manveru/faker"
)

// Seeder contains the methods needed to seed the app
type Seeder struct {
	bluebox.HTTPService
	dbclient *db.DB
	faker    *faker.Faker
}

// New returns a new Seeder
func New(httpService bluebox.HTTPService, dbinstance *db.DB) *Seeder {
	f, err := faker.New("en")
	if err != nil {
		panic(err)
	}
	return &Seeder{
		HTTPService: httpService,
		dbclient:    dbinstance,
		faker:       f,
	}
}

// Seed seeds the database with an admin user and runs the seeder function.
func (s *Seeder) Seed() error {
	// Create the admin
	email := "admin@example.com"
	s.dbclient.UserService().UserCreate(&bluebox.User{
		FirstName:    s.faker.FirstName(),
		LastName:     s.faker.LastName(),
		SessionToken: helpers.NewSessionToken(),
		Email:        email,
		Password:     helpers.HashPassword("password"),
		Role:         bluebox.RoleAdmin,
		Archived:     false,
		Disabled:     false,
		Data:         &bluebox.UserData{},
	})
	return s.Run()
}

// Run will begin the seeding process. The server must be running.
func (s *Seeder) Run() error {
	// Login first
	err := s.HTTPService.UserSignIn(&bluebox.UserSignInRequest{
		Email:    "admin@example.com",
		Password: "password",
	})
	if err != nil {
		return err
	}

	err = s.SeedUsers()
	if err != nil {
		return err
	}

	log.Println("Seeding completed without errors.")
	return nil
}

// SeedUsers will create some users
func (s *Seeder) SeedUsers() error {
	log.Println("Seeding users...")

	// Create some staff users
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			firstName := s.faker.FirstName()
			email := fmt.Sprintf("%s@example.com", strings.ToLower(firstName))
			newUser := &bluebox.UserCreateRequest{
				FirstName: firstName,
				LastName:  s.faker.LastName(),
				Email:     email,
				Password:  "password",
				Role:      bluebox.RoleStaff,
				Data:      &bluebox.UserData{},
			}
			err := s.HTTPService.UserCreate(newUser)
			if err != nil {
				panic(err)
			}
			log.Println("New user created:", newUser.FirstName)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}
