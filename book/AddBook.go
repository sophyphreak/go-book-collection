package book

// AddBook adds a new book to Collection
func (c *Collection) AddBook(title, author, publisher string, pages, publicationYear int) int {
	var id int
	if len(c.Collection) == 0 {
		id = 0
	} else {
		id = c.Collection[len(c.Collection)-1].ID + 1
	}
	newBook := Book{
		ID:              id,
		Title:           title,
		Author:          author,
		Publisher:       publisher,
		Pages:           pages,
		PublicationYear: publicationYear,
	}
	c.Collection = append(c.Collection, newBook)
	return id
}
