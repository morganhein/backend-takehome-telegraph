package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph/store"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph/transport"
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
	m := mux.NewRouter()
	m.HandleFunc("/equipment", srv.ListEquipment)
	m.HandleFunc("/events", srv.ListEvents)
	m.HandleFunc("/locations", srv.ListLocations)
	m.HandleFunc("/waybills/", srv.ListWaybills)
	m.HandleFunc("/waybills/{waybillID}", srv.Getwaybill)
	m.HandleFunc("/waybills/{waybillID}/{association}", srv.Getwaybill)

	// start
	if err := http.ListenAndServe(":8080", m); err != nil {
		panic(err)
	}
}
