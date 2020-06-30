package book

import "testing"

func TestDeleteBook(t *testing.T) {
	var c Collection
	book := Book{0, "The Alchemist", "Paulo Coelho", "Some publisher", 200, 1989}
	c.Collection = append(c.Collection, book)
	c.DeleteBook(0)
	if len(c.Collection) != 0 {
		t.Errorf("Delete book action was not successful")
	}
	err := c.DeleteBook(-1)
	if err == nil {
		t.Errorf("Delete book with index out of bounds succeeded when it should've failed")
	}
}
