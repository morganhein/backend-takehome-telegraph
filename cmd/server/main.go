package main

import (
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph/store"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph/transport"
	"net/http"
)

var (
	connString = "postgres://candidate:password123@localhost:5432/takehome?sslmode=disable"
)

func main() {
	// instantiation
	db, err := store.CreatePostgresStore(connString)
	if err != nil {
		panic(err)
	}
	srv := transport.NewHTTPTransport(telegraph.New(db))

	// muxing
	mux := http.NewServeMux()
	mux.HandleFunc("/equipment", srv.ListEquipment)
	mux.HandleFunc("/events", srv.ListEvents)
	mux.HandleFunc("/locations", srv.ListLocations)
	mux.HandleFunc("/waybills", srv.ListWaybills)

	// start
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
