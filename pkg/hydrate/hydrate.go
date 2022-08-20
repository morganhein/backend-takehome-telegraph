package hydrate

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph/store"
)

type models interface {
	telegraph.Equipment | telegraph.Event | telegraph.Location | telegraph.Waybill
}

func Start(pgString string) error {
	//migrate first, and ready the db
	if err := prepare(pgString); err != nil {
		return err
	}
	st, _, err := store.CreatePostgresStore(pgString)
	if err != nil {
		return err
	}
	if err = process[telegraph.Equipment](st.CreateEquipment, "data/equipment.csv"); err != nil {
		return err
	}
	if err = process[telegraph.Event](st.CreateEvent, "data/events.csv"); err != nil {
		return err
	}
	if err = process[telegraph.Location](st.CreateLocation, "data/locations.csv"); err != nil {
		return err
	}
	return nil
}

func prepare(pgString string) error {
	m, err := migrate.New(
		"file://db/migrations",
		pgString)
	if err != nil {
		return err
	}
	err = m.Down()
	if err != nil {
		return err
	}
	return m.Up()
}

func process[T models](f func(T) error, fileLoc string) error {
	x, err := loadFile[T](fileLoc)
	if err != nil {
		panic(err)
	}
	for _, v := range x {
		err = f(v)
		if err != nil {
			panic(err)
		}
	}
	return nil
}

func loadFile[T interface{}](f string) ([]T, error) {
	clientsFile, err := os.OpenFile(f, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = clientsFile.Close()
	}()

	entries := []T{}

	if err := gocsv.UnmarshalFile(clientsFile, &entries); err != nil {
		panic(err)
	}

	return entries, nil
}
