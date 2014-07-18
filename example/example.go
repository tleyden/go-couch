package main

import (
	"io"

	"github.com/couchbaselabs/logg"
	couch "github.com/tleyden/go-couch"
)

func main() {

	err := dumpChangesFeed("http://localhost:5984/test")
	if err != nil {
		logg.LogError(err)
		panic("Error dumping changes feed")
	}

}

func dumpChangesFeed(dburl string) error {
	db, err := couch.Connect(dburl)
	if err != nil {
		return err
	}
	logg.Log("connected to %v", db)
	options := map[string]interface{}{}
	options["since"] = nil
	err = db.Changes(
		func(reader io.Reader) interface{} {
			logg.Log("changes feed called with %v", reader)
			changes, err := couch.ReadAllChanges(reader)
			if err != nil {
				logg.Log("Error reading changes %v", err)
			}
			dumpChanges(changes)
			return nil
		}, options)
	return nil
}

func dumpChanges(changes couch.Changes) {
	logg.Log("Changes up to last sequence: %v", changes.LastSequence)
	for _, change := range changes.Results {
		logg.Log("Id: %v Deleted: %v", change.Id, change.Deleted)
	}
}
