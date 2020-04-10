package book_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentbookef/domain"
	"testing"
	"time"
)

/***
HOW IT WORKS
We Testing the create method
1- we create an Object
2- we sending the object to the method CreateBookImage()
2- if all good the method CreateBookImage(), will send the Object to the Backend
3- if all good the method CreateBookImage(), will return expected response in result and nil in err variables
3- If an error occurs the method CreateBookImage() will return an error message in both variables(result and err)
*/
/***
BOOK IMAGE TEST START HERE
*/
func TestCreatBookImage(t *testing.T) {
	bImage := domain.BookImage{"000", "0000", "test"}
	result, err := CreatBookImage(bImage)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBookImage(t *testing.T) {
	result, err := ReadBookImage("")
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBookImages(t *testing.T) {
	result, err := ReadBookImages()
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestUpdateBookImage(t *testing.T) {
	bImage := domain.BookImage{"000", "0000", "test"}
	result, err := UpdateBookImage(bImage)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestDeleteBookImage(t *testing.T) {
	bImage := domain.BookImage{"000", "0000", "test"}
	result, err := DeleteBookImage(bImage)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}

/***
BOOK TEST IMAGE ENDS HERE
*/

/***
BOOK TEST START HERE
*/
func TestCreateBook(t *testing.T) {
	book := domain.Book{"0000", "biblia", "ENglish", "VI", 300}
	result, err := CreateBook(book)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBook(t *testing.T) {
	result, err := ReadBook("")
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBooks(t *testing.T) {
	result, err := ReadBooks()
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestDeleteBook(t *testing.T) {
	book := domain.Book{"0000", "biblia", "ENglish", "VI", 300}
	result, err := DeleteBook(book)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestUpdateBook(t *testing.T) {
	book := domain.Book{"0000", "biblia", "ENglish", "VI", 300}
	result, err := UpdateBook(book)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}

/***
BOOK TEST ENdS HERE
*/

/***
BOOKDEPARTMENT TEST START HERE
*/
func TestCreateBookdepartment(t *testing.T) {
	bookDeparment := domain.BookDepartment{"0000", "00000", "test"}
	result, err := CreateBookdepartment(bookDeparment)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestDeleteBookDepartment(t *testing.T) {
	bookDeparment := domain.BookDepartment{"0000", "00000", "test"}
	result, err := DeleteBookDepartment(bookDeparment)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestUpdateBookDepartment(t *testing.T) {
	bookDeparment := domain.BookDepartment{"0000", "00000", "test"}
	result, err := UpdateBookDepartment(bookDeparment)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBookDepartment(t *testing.T) {
	result, err := ReadBookDepartment("")
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBookDepartments(t *testing.T) {
	result, err := ReadBookDepartments()
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}

/***
BOOKDEPARTMENT TEST ENDS HERE
*/

/***
BOOKPOST TEST START HERE
*/
func TestCreatBookPost(t *testing.T) {
	post := domain.BookPost{"000", "espoirditekemena@gmail.com", "0000", time.Now(), "00034", "on", "all the page od thid book are in good sharp"}
	result, err := CreatBookPost(post)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestUpdateBookPost(t *testing.T) {
	post := domain.BookPost{"000", "espoirditekemena@gmail.com", "0000", time.Now(), "00034", "on", "all the page od thid book are in good sharp"}
	result, err := UpdateBookPost(post)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBookPosts(t *testing.T) {
	result, err := ReadBookPosts()
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestDeleteBookPost(t *testing.T) {
	post := domain.BookPost{"000", "espoirditekemena@gmail.com", "0000", time.Now(), "00034", "on", "all the page od thid book are in good sharp"}
	result, err := DeleteBookPost(post)
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
func TestReadBookPost(t *testing.T) {
	result, err := ReadBookPost("")
	assert.Nil(t, err)
	fmt.Println("result is: ", result)
}
