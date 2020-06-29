package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"./book"
)

func populateCollection(collection *book.Collection) {
	jsonFile, err := os.Open("collection.json")
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(byteValue, &collection)
}
