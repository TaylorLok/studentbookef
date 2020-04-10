package book_io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const bookURL = api.BASE_URL + "book/"

func CreateBook(mybook domain.Book) (domain.Book, error) {
	entity := domain.Book{}
	resp, _ := api.Rest().SetBody(mybook).Post(bookURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBook(id string) (domain.Book, error) {
	entity := domain.Book{}
	resp, _ := api.Rest().Get(bookURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBooks() ([]domain.Book, error) {
	entity := []domain.Book{}
	resp, _ := api.Rest().Get(bookURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBook(mybook domain.Book) (domain.Book, error) {
	entity := domain.Book{}
	resp, _ := api.Rest().SetBody(mybook).Post(bookURL + "detele")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBook(mybook domain.Book) (domain.Book, error) {
	entity := domain.Book{}
	resp, _ := api.Rest().SetBody(mybook).Post(bookURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
