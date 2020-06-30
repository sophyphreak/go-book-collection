package book

import (
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
}
