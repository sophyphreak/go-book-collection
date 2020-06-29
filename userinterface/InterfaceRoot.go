package userinterface

import (
	"fmt"
	"strings"

	"../book"
)

// InterfaceRoot is the root of the user interface
func InterfaceRoot(c book.Collection) {
	fmt.Println("Welcome to your book collection")
	for {
		var userInput string
		fmt.Println("Quit by inputting \"q\"")
		fmt.Print("Would you like to display your book collection, add a new book, update a book, or delete a book? ")
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
		if userInput == "q" {
			break
		}
		switch userInput {
		case "display", "dis", "di":
		case "add", "ad", "a":
		case "update", "up", "u":
		case "delete", "del", "de":
		default:
			fmt.Println("Sorry, I did not understand that input. Please choose one of the given options.")
			continue
		}
		fmt.Println("success")
	}
	fmt.Println("Thank you for using your book collection. Everything is saved for next time! :-)")
}
