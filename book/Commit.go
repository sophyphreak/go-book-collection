package book

import (
	"encoding/json"
	"io/ioutil"
)

// Commit saves the current collection to the collection.json file
func (c *Collection) Commit() {
	f, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("collection.json", f, 0644)
}
