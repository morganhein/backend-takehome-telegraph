package hydrate

import (
	"github.com/jackc/pgx/v4"
	"os"

	"github.com/gocarina/gocsv"

	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph"
	"github.com/morganhein/backend-takehome-telegraph/internal/telegraph/store"
)

type models interface {
	telegraph.Equipment | telegraph.Event | telegraph.Location | telegraph.Waybill
}

func Start(pgString string) error {
	//migrate first, and ready the db

	st, _, err := store.CreatePostgresStore(pgString)
	if err != nil {
		panic(err)
	}
	err = process[telegraph.Equipment](st.CreateEquipment, "data/equipment.csv")
	if err != nil {
		return err
	}
	err = process[telegraph.Event](st.CreateEvent, "data/events.csv")
	if err != nil {
		return err
	}
	return nil
}

func migrate(db *pgx.Conn) error {
	return nil
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
