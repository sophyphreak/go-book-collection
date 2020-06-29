package book

// DeleteBook deletes a book from Book
func (c *Collection) DeleteBook(id int) error {
	index, _, err := c.findByID(id)
	if err != nil {
		return err
	}
	c.Collection[len(c.Collection)-1], c.Collection[index] = c.Collection[index], c.Collection[len(c.Collection)-1]
	c.Collection = c.Collection[:len(c.Collection)-1]
	c.Commit()
	return nil
}
