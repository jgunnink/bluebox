package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jgunnink/bluebox"
	"github.com/jgunnink/bluebox/client"
	"github.com/jgunnink/bluebox/db"
	"github.com/jgunnink/bluebox/seed"
	"github.com/jgunnink/bluebox/server"
)

var helptext = `
Please use one of the following commands:
	serve
	db:drop
	db:migrate
	db:seed
Have a pleasant day!
`

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(helptext)
		os.Exit(1)
	}

	db := db.New(
		bluebox.DBname,
		bluebox.DBuser,
		bluebox.DBpass,
		bluebox.DBhost,
		bluebox.DBport,
	)

	defer db.Close()

	c := client.New("http://localhost:8080")

	for _, arg := range args {
		fmt.Println("Running command:", arg)
		switch arg {
		case "serve":
			serve(db)
		case "db:migrate":
			db.Migrate()
		case "db:drop":
			db.Drop()
		case "db:seed":
			seeder(c, db)
		case "db:setup":
			db.Drop()
			db.Migrate()
			seeder(c, db)
		default:
			fmt.Printf("%s is not a valid command\n", arg)
			fmt.Println(helptext)
		}
	}
}

func serve(db *db.DB) {
	s := server.New(":8080", db.AuthService(), db.UserService())
	log.Fatalln(s.Run())
}

func seeder(c bluebox.HTTPService, db *db.DB) {
	s := seed.New(c, db)
	err := s.Seed()
	if err != nil {
		log.Println(err)
		return
	}
}
