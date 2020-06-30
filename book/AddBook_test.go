package book

import (
	"os"
	"testing"
)

func TestAddBook(t *testing.T) {
	var c Collection
	c.Populate()

	title := "The Alchemist"
	author := "Paulo Coelho"
	publisher := "Some publisher"
	pages := 200
	publicationYear := 1989
	c.AddBook(title, author, publisher, pages, publicationYear)

	book := c.Collection[len(c.Collection)-1]
	if book.Title != title {
		t.Errorf("book.Title should be %s instead got %s", title, book.Title)
	}
	if book.Author != author {
		t.Errorf("book.Author should be %s instead got %s", author, book.Author)
	}
	if book.Publisher != publisher {
		t.Errorf("book.Publisher should be %s instead got %s", publisher, book.Publisher)
	}
	if book.Pages != pages {
		t.Errorf("book.Pages should be %d instead got %d", pages, book.Pages)
	}
	if book.PublicationYear != publicationYear {
		t.Errorf("book.PublicationYear should be %d instead got %d", publicationYear, book.PublicationYear)
	}
	os.Remove("collection.json")
}
