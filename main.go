package main

import (
	"encoding/json"
	"io/ioutil"

	"./books"
)

func main() {
	var collection books.Books
	populateCollection(&collection)

	collection.AddBook("The Alchemist", "Paulo Coelho", "Some Publisher", 200, 1989)

	f, err := json.MarshalIndent(collection, "", "\t")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("books.json", f, 0644)
}
