package books

// DeleteBook deletes a book from Book
func (b *Books) DeleteBook(id int) {
	index, _ := b.findByID(id)
	b.Books[len(b.Books)-1], b.Books[index] = b.Books[index], b.Books[len(b.Books)-1]
}
