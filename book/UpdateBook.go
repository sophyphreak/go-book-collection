package book

// UpdateTitle updates a book's title
func (c *Collection) UpdateTitle(id int, title string) {
	bookIndex, book := c.findByID(id)
	book.Title = title
	c.Collection[bookIndex] = book
}

// UpdateAuthor updates a book's author
func (c *Collection) UpdateAuthor(id int, author string) {
	bookIndex, book := c.findByID(id)
	book.Author = author
	c.Collection[bookIndex] = book
}

// UpdatePublisher updates a book's publisher
func (c *Collection) UpdatePublisher(id int, publisher string) {
	bookIndex, book := c.findByID(id)
	book.Publisher = publisher
	c.Collection[bookIndex] = book
}

// UpdatePages updates a book's pages
func (c *Collection) UpdatePages(id, pages int) {
	bookIndex, book := c.findByID(id)
	book.Pages = pages
	c.Collection[bookIndex] = book
}

// UpdatePublicationYear updates a book's publicationYear
func (c *Collection) UpdatePublicationYear(id, publicationYear int) {
	bookIndex, book := c.findByID(id)
	book.PublicationYear = publicationYear
	c.Collection[bookIndex] = book
}
