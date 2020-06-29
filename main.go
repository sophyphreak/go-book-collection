package main

import (
	"./book"
)

func main() {
	var collection book.Collection
	collection.Populate()
	collection.AddBook("The Alchemist", "Paulo Coelho", "Some Publisher", 200, 1989)
}
