package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"../book"
)

func addBook(c *book.Collection) {
	var userInput, title, author, publisher string
	var pages, publicationYear int
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("I see you want to add a book.")
	fmt.Print("What is the title? ")
	scanner.Scan()
	title = scanner.Text()
	fmt.Print("Who is the author? ")
	scanner.Scan()
	author = scanner.Text()
	fmt.Print("Who is the publisher? ")
	scanner.Scan()
	publisher = scanner.Text()
	for {
		fmt.Print("How many pages in the book? ")
		fmt.Scanln(&userInput)
		pages, err = strconv.Atoi(userInput)
		if err != nil || pages < 0 {
			fmt.Println("Please enter a whole number")
			continue
		}
		break
	}
	for {
		fmt.Print("What is the publication year? ")
		fmt.Scanln(&userInput)
		publicationYear, err = strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter an integer")
			continue
		}
		break
	}
	c.AddBook(title, author, publisher, pages, publicationYear)
}
