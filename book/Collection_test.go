package book

import (
	"testing"
)

func TestFindByID(t *testing.T) {
	var c Collection
	book1 := Book{0, "The Alchemist", "Paulo Coelho", "Some publisher", 200, 1989}
	c.Collection = append(c.Collection, book1)
	i, foundBook, err := c.findByID(0)
	if err != nil {
		t.Errorf("Incorrectly recieved error message from function. Error message text: %s.", err)
	}
	if foundBook != book1 {
		t.Errorf("Return value should be %+v but instead recieved %+v", book1, foundBook)
	}
	if i != 0 {
		t.Errorf("Returned index should be 0 but instead recieved %d", i)
	}

	book2 := Book{1, "Zhuangzi", "Zhuangzi", "Some publisher", 500, -400}
	c.Collection = append(c.Collection, book2)
	i, foundBook, err = c.findByID(0)
	if err != nil {
		t.Errorf("Incorrectly recieved error message from function. Error message text: %s.", err)
	}
	if foundBook != book1 {
		t.Errorf("Return value should be %+v but instead recieved %+v", book1, foundBook)
	}
	if i != 0 {
		t.Errorf("Returned index should be 0 but instead recieved %d", i)
	}

	i, foundBook, err = c.findByID(1)
	if err != nil {
		t.Errorf("Incorrectly recieved error message from function. Error message text: %s.", err)
	}
	if foundBook != book2 {
		t.Errorf("Return value should be %+v but instead recieved %+v", book2, foundBook)
	}
	if i != 1 {
		t.Errorf("Returned index should be 1 but instead recieved %d", i)
	}

	i, foundBook, err = c.findByID(-1)
	if err == nil {
		t.Errorf("Incorrectly recieved nil error message")
	}
	if (foundBook != Book{}) {
		t.Errorf("Return value should be %+v but instead recieved %+v", Book{}, foundBook)
	}
	if i != 0 {
		t.Errorf("Returned index should be 0 but instead recieved %d", i)
	}
}
