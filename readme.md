# Bluebox

A back-end webservice in Go which can be used as a template or starting point to get going quickly.

## Developer

[JK Gunnink](mailto:jgunnink@gmail.com)

## What's included

*   Users can signup for an account
*   Roles are assigned to users (Admin, User)
*   Admins can manage users
*   Admins can send password reset emails to users - a sendgrid API key is required (specify this in the go server package)
*   Users can reset their passwords with a unique password reset link which is emailed to them
*   Admins can disable user accounts, preventing them from logging in.

## Getting Started

### Clone the project repo

```bash
git clone git@github.com:jgunnink/bluebox.git $GOPATH/src/github.com/jgunnink/bluebox
cd $_
```

### Install go dependencies

```bash
dep ensure
```

### Start services

1.  Ensure postgres is listening for connections. A docker example: `docker run --rm -p 5432:5432 --name db -e POSTGRES_USER=develop -e POSTGRES_PASSWORD=develop -e POSTGRES_DB=bluebox postgres`
1.  `dep ensure`
1.  Start the go server with `go run ./cmd/bluebox/main.go serve`
1.  Seed the database with `go run cmd/bluebox/*.go db:setup`

## Dependencies

1.  PostgreSQL
1.  Golang (ensure you run dep ensure to fetch imported packages)

## Notes

It is intended that this Go backend server (bluebox) is to be used in tandem with a SPA (i.e. React or Vue), as such the Go server runs on port 8000, and the react webserver runs on port 3000. The Go server reverse proxies the react server. Accessing the application on port 3000 in your browser will cause unhappiness :smiley:
