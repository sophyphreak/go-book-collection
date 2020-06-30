package main

import (
	"./book"
	"./userinterface"
)

func main() {
	var collection book.Collection
	collection.Populate()
	userinterface.InterfaceRoot(&collection)
}
