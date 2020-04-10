package domain

import "time"

type Book struct {
	Id       string  `json:"id"`
	Title    string  `json:"title"`
	Language string  `json:"language"`
	Edition  string  `json:"edition"`
	Price    float64 `json:"price"`
}

type BookDepartment struct {
	BookId       string `json:"book_Id"`
	DepartmentId string `json:"department_Id"`
	Description  string `json:"description"`
}
type BookImage struct {
	BookId      string `json:"book_id"`
	ImageId     string `json:"image_id"`
	Description string `json:"description"`
}
type BookPost struct {
	Id          string    `json:"id"`
	Email       string    `json:"email"`
	BookId      string    `json:"book_Id"`
	Date        time.Time `json:"date"`
	LocationId  string    `json:"location_Id"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}
