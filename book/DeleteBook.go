package book

// DeleteBook deletes a book from Book
func (c *Collection) DeleteBook(id int) {
	index, _ := c.findByID(id)
	c.Collection[len(c.Collection)-1], c.Collection[index] = c.Collection[index], c.Collection[len(c.Collection)-1]
	c.Commit()
}
