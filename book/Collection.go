package book

// Collection is an array of Book
type Collection struct {
	Collection []Book `json:"collection"`
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

func (c *Collection) findByID(id int) (int, Book) {
	var bookIndex int
	var book Book
	for i, v := range c.Collection {
		if v.ID == id {
			bookIndex = i
			book = v
		}
	}
	return bookIndex, book
}
