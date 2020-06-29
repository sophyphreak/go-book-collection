package userinterface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../book"
)

func updateBook(c *book.Collection) {
	var userInput string
	var id int
	var err error
	for {
		fmt.Print("What is the id of the book you want to update? ")
		fmt.Scanln(&userInput)
		id, err = strconv.Atoi(userInput)
		if err != nil || id < 0 {
			fmt.Println("Please provide a whole number.")
			continue
		}
		break
	}
	for {
		fmt.Print("Update the title, author, publisher, pages, or publication year? ")
		fmt.Scanln(&userInput)
		userInput = strings.ToLower(userInput)
		switch userInput {
		case "title", "titl", "tit", "ti", "t":
			updateTitle(c, id)
		case "author", "autho", "auth", "aut", "au", "a":
			updateAuthor(c, id)
		case "publisher", "publish":
			updatePublisher(c, id)
		case "pages", "page", "pag", "pg":
			updatePages(c, id)
		case "publication", "public", "publicatio", "publicati", "publicat", "publica":
			updatePublicationYear(c, id)
		default:
			fmt.Println("Sorry, I don't understand that.")
			continue
		}
		break
	}
}

func updateTitle(c *book.Collection, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the new title? ")
	scanner.Scan()
	title := scanner.Text()
	err := c.UpdateTitle(id, title)
	if err != nil {
		fmt.Println("Could not update book.")
		fmt.Println(err)
	}
}

func updateAuthor(c *book.Collection, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the new author? ")
	scanner.Scan()
	author := scanner.Text()
	err := c.UpdateAuthor(id, author)
	if err != nil {
		fmt.Println("Could not update book.")
		fmt.Println(err)
	}
}

func updatePublisher(c *book.Collection, id int) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("What is the new publisher? ")
	scanner.Scan()
	publisher := scanner.Text()
	err := c.UpdatePublisher(id, publisher)
	if err != nil {
		fmt.Println("Could not update book.")
		fmt.Println(err)
	}
}

func updatePages(c *book.Collection, id int) {
	var userInput string
	var pages int
	var err error
	for {
		fmt.Print("What is the new number of pages? ")
		fmt.Scanln(&userInput)
		pages, err = strconv.Atoi(userInput)
		if err != nil || pages < 0 {
			fmt.Println("Please input a positive integer.")
		}
		break
	}
	err = c.UpdatePages(id, pages)
	if err != nil {
		fmt.Println("Could not update book.")
		fmt.Println(err)
	}
}

func updatePublicationYear(c *book.Collection, id int) {
	var userInput string
	var publicationYear int
	var err error
	for {
		fmt.Print("What is the new publication year? ")
		fmt.Scanln(&userInput)
		publicationYear, err = strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please input an integer.")
		}
		break
	}
	err = c.UpdatePublicationYear(id, publicationYear)
	if err != nil {
		fmt.Println("Could not update book.")
		fmt.Println(err)
	}
}
