package userinterface

import (
	"fmt"
	"strconv"

	"../book"
)

func deleteBook(c *book.Collection) {
	var userInput string
	for {
		fmt.Print("What is the id of the book you want to delete? ")
		fmt.Scanln(&userInput)
		id, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please provide a valid integer.")
			continue
		}
		err = c.DeleteBook(id)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Book with id of", id, "successfully deleted.")
		}
		break
	}
}
