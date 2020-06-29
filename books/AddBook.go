package books

// AddBook adds a new book to Books
func (b *Books) AddBook(title, author, publisher string, pages, publicationYear int) int {
	var id int
	if len(b.Books) == 0 {
		id = 0
	} else {
		id = b.Books[len(b.Books)-1].ID + 1
	}
	newBook := Book{
		ID:              id,
		Title:           title,
		Author:          author,
		Publisher:       publisher,
		Pages:           pages,
		PublicationYear: publicationYear,
	}
	b.Books = append(b.Books, newBook)
	return id
}
