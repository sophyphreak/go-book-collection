package book

import "testing"

func TestUpdateTitle(t *testing.T) {
	var c Collection
	book := Book{0, "Alchemist", "Paulo Coelho", "Some publisher", 200, 1989}
	c.Collection = append(c.Collection, book)
	err := c.UpdateTitle(0, "The Alchemist")
	if err != nil {
		t.Errorf("Tried to update title but recieved error: %s", err)
	}
	title := c.Collection[0].Title
	if title != "The Alchemist" {
		t.Errorf("Tried to update title to \"The Alchemist\" but instead recieved %s", title)
	}

	err = c.UpdateTitle(-1, "anything")
	if err == nil {
		t.Errorf("Expected an error but instead recieved nil")
	}
}

func TestUpdateAuthor(t *testing.T) {
	var c Collection
	book := Book{0, "The Alchemist", "Paulo", "Some publisher", 200, 1989}
	c.Collection = append(c.Collection, book)
	err := c.UpdateAuthor(0, "Paulo Coelho")
	if err != nil {
		t.Errorf("Tried to update author but recieved error: %s", err)
	}
	author := c.Collection[0].Author
	if author != "Paulo Coelho" {
		t.Errorf("Tried to update title to \"Paulo Coelho\" but instead recieved %s", author)
	}

	err = c.UpdateAuthor(-1, "anything")
	if err == nil {
		t.Errorf("Expected an error but instead recieved nil")
	}
}

func TestUpdatePublisher(t *testing.T) {
	var c Collection
	book := Book{0, "The Alchemist", "Paulo Coelho", "Some", 200, 1989}
	c.Collection = append(c.Collection, book)
	err := c.UpdatePublisher(0, "Some publisher")
	if err != nil {
		t.Errorf("Tried to update publisher but recieved error: %s", err)
	}
	publisher := c.Collection[0].Publisher
	if publisher != "Some publisher" {
		t.Errorf("Tried to update publisher to \"Some publisher\" but instead recieved %s", publisher)
	}

	err = c.UpdatePublisher(-1, "anything")
	if err == nil {
		t.Errorf("Expected an error but instead recieved nil")
	}
}

func TestUpdatePages(t *testing.T) {
	var c Collection
	book := Book{0, "The Alchemist", "Paulo Coelho", "Some publisher", 20, 1989}
	c.Collection = append(c.Collection, book)
	err := c.UpdatePages(0, 200)
	if err != nil {
		t.Errorf("Tried to update pages but recieved error: %s", err)
	}
	pages := c.Collection[0].Pages
	if pages != 200 {
		t.Errorf("Tried to update pages to 200 but instead recieved %d", pages)
	}

	err = c.UpdatePages(-1, 200)
	if err == nil {
		t.Errorf("Expected an error but instead recieved nil")
	}
}

func TestUpdatePublicationYear(t *testing.T) {
	var c Collection
	book := Book{0, "The Alchemist", "Paulo Coelho", "Some publisher", 200, 198}
	c.Collection = append(c.Collection, book)
	err := c.UpdatePublicationYear(0, 1989)
	if err != nil {
		t.Errorf("Tried to update publicationYear but recieved error: %s", err)
	}
	publicationYear := c.Collection[0].PublicationYear
	if publicationYear != 1989 {
		t.Errorf("Tried to update publicationYear to 1989 but instead recieved %d", publicationYear)
	}

	err = c.UpdatePublicationYear(-1, 1989)
	if err == nil {
		t.Errorf("Expected an error but instead recieved nil")
	}
}
