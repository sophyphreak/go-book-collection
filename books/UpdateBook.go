package books

// UpdateTitle updates a book's title
func (b *Books) UpdateTitle(id int, title string) {
	bookIndex, book := b.findByID(id)
	book.Title = title
	b.Books[bookIndex] = book
}

// UpdateAuthor updates a book's author
func (b *Books) UpdateAuthor(id int, author string) {
	bookIndex, book := b.findByID(id)
	book.Author = author
	b.Books[bookIndex] = book
}

// UpdatePublisher updates a book's publisher
func (b *Books) UpdatePublisher(id int, publisher string) {
	bookIndex, book := b.findByID(id)
	book.Publisher = publisher
	b.Books[bookIndex] = book
}

// UpdatePages updates a book's pages
func (b *Books) UpdatePages(id, pages int) {
	bookIndex, book := b.findByID(id)
	book.Pages = pages
	b.Books[bookIndex] = book
}

// UpdatePublicationYear updates a book's publicationYear
func (b *Books) UpdatePublicationYear(id, publicationYear int) {
	bookIndex, book := b.findByID(id)
	book.PublicationYear = publicationYear
	b.Books[bookIndex] = book
}
