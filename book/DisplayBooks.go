package book

import "fmt"

// DisplayCollection prints all data from all books
func (c *Collection) DisplayCollection() {
	if len(c.Collection) == 0 {
		fmt.Println()
		fmt.Println("No books in collection.")
		fmt.Println()
		return
	}
	for _, v := range c.Collection {
		fmt.Println()
		fmt.Println("id:", v.ID)
		fmt.Println("title:", v.Title)
		fmt.Println("author:", v.Author)
		fmt.Println("publisher:", v.Publisher)
		fmt.Println("pages:", v.Pages)
		fmt.Println("publication year:", v.PublicationYear)
	}
	fmt.Println()
}
