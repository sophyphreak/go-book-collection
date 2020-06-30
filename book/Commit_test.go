package book

import (
	"os"
	"testing"
)

func TestCommit(t *testing.T) {
	var c Collection
	book := Book{0, "The Alchemist", "Paulo Coelho", "Some publisher", 200, 1989}
	c.Collection = append(c.Collection, book)
	err := c.Commit()
	if err != nil {
		t.Errorf("Error when attempting commit %s", err)
	}
	var d Collection
	d.Populate()
	if len(d.Collection) == 0 {
		t.Errorf("Retrieved collection is empty. Expecting: %+v", c)
	}
	for i, v := range c.Collection {
		if v != d.Collection[i] {
			t.Errorf("Expected to retrieve %+v from file but instead recieved %+v", c, d)
		}
	}
	os.Remove("collection.json")
}
