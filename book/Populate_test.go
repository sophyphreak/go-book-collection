package book

import (
	"os"
	"testing"
)

func TestPopulate(t *testing.T) {
	defer func() {
		if recover() != nil {
			t.Errorf("Populate failed")
		}
	}()

	var c Collection
	c.Populate()
	os.Remove("collection.json")
}
