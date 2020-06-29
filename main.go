package main

import (
	"encoding/json"
	"io/ioutil"

	"./book"
)

func main() {
	var collection book.Collection
	collection.Populate()

	collection.AddBook("The Alchemist", "Paulo Coelho", "Some Publisher", 200, 1989)

	f, err := json.MarshalIndent(collection, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("collection.json", f, 0644)
}
