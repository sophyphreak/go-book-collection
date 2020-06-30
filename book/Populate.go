package book

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// Populate populates a given collection with the json data in collection.json
func (c *Collection) Populate() {
	if _, err := os.Stat("collection.json"); os.IsNotExist(err) {
		os.Create("collection.json")
	}
	jsonFile, err := os.Open("collection.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, c)
}
