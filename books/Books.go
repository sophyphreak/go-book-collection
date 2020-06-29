package books

// Books is an array of Book
type Books struct {
	Books []Book `json:"books"`
}

// Book is a single book
type Book struct {
	ID              int    `json:"id"`
	Title           string `json:"title"`
	Author          string `json:"author"`
	Publisher       string `json:"publisher"`
	Pages           int    `json:"pages"`
	PublicationYear int    `json:"publicationYear"`
}

func (b *Books) findByID(id int) (int, Book) {
	var bookIndex int
	var book Book
	for i, v := range b.Books {
		if v.ID == id {
			bookIndex = i
			book = v
		}
	}
	return bookIndex, book
}
