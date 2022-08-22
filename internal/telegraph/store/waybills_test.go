package store

import "testing"

var (
	connString = "postgres://candidate:password123@localhost:5432/takehome?sslmode=disable"
)

func TestPostgres_GetWaybill_WithAssociation(t *testing.T) {
	db, err := CreatePostgresStore(connString)
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}
	res, err := db.GetWaybill("6", "equipment")
	if err != nil {
		t.Log(err)
		t.FailNow()
		return
	}
	t.Log(res)
}
