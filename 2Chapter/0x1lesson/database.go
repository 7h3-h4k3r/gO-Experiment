package main

import "fmt"

type Datastore interface {
	Save(data string) error
}

func process(ds Datastore) error {
	return ds.Save("importand data")
}

type DB struct{}

func (db DB) Save(data string) error {
	fmt.Println("Saving data to DB :", data)
	return nil
}

func main() {
	database := DB{}
	process(database)
}
