package book

import "fmt"

// DisplayCollection prints all data from all books
func (c *Collection) DisplayCollection() {
	for i, v := range c.Collection {
		fmt.Println("")
		fmt.Println("Book", i+1)
		fmt.Println("id:", v.ID)
		fmt.Println("title:", v.Title)
		fmt.Println("author:", v.Author)
		fmt.Println("publisher:", v.Publisher)
		fmt.Println("pages:", v.Pages)
		fmt.Println("publication year:", v.PublicationYear)
	}
}
