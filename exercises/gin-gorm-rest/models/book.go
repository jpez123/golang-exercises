package models

//Book - Backticks are used to tag each field from database
type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

//CreateBookInput - validate user's input
type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

//UpdateBookInput - updates book
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}
