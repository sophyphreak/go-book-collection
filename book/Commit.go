package book

import (
	"encoding/json"
	"io/ioutil"
)

// Commit saves the current collection to the collection.json file
func (c *Collection) Commit() error {
	f, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("collection.json", f, 0644)
	return err
}
