package main

import (
	"./book"
	"./userinterface"
)

func main() {
	var collection book.Collection
	collection.Populate()
	userinterface.InterfaceRoot(&collection)
	// collection.AddBook("The Alchemist", "Paulo Coelho", "Some Publisher", 200, 1989)
}
