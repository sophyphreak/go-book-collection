package book

// UpdateTitle updates a book's title
func (c *Collection) UpdateTitle(id int, title string) error {
	bookIndex, book, err := c.findByID(id)
	if err != nil {
		return err
	}
	book.Title = title
	c.Collection[bookIndex] = book
	c.Commit()
	return nil
}

// UpdateAuthor updates a book's author
func (c *Collection) UpdateAuthor(id int, author string) error {
	bookIndex, book, err := c.findByID(id)
	if err != nil {
		return err
	}
	book.Author = author
	c.Collection[bookIndex] = book
	c.Commit()
	return nil
}

// UpdatePublisher updates a book's publisher
func (c *Collection) UpdatePublisher(id int, publisher string) error {
	bookIndex, book, err := c.findByID(id)
	if err != nil {
		return err
	}
	book.Publisher = publisher
	c.Collection[bookIndex] = book
	c.Commit()
	return nil
}

// UpdatePages updates a book's pages
func (c *Collection) UpdatePages(id, pages int) error {
	bookIndex, book, err := c.findByID(id)
	if err != nil {
		return err
	}
	book.Pages = pages
	c.Collection[bookIndex] = book
	c.Commit()
	return nil
}

// UpdatePublicationYear updates a book's publicationYear
func (c *Collection) UpdatePublicationYear(id, publicationYear int) error {
	bookIndex, book, err := c.findByID(id)
	if err != nil {
		return err
	}
	book.PublicationYear = publicationYear
	c.Collection[bookIndex] = book
	c.Commit()
	return nil
}
