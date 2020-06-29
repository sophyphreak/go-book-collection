package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"./books"
)

func populateCollection(collection *books.Books) {
	jsonFile, err := os.Open("books.json")
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
