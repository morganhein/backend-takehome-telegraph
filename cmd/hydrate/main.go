package main

import "github.com/morganhein/backend-takehome-telegraph/pkg/hydrate"

var (
	connString = "postgres://candidate:password123@localhost:5432/takehome?sslmode=disable"
)

func main() {
	err := hydrate.Start(connString)
	if err != nil {
		panic(err)
	}
}
