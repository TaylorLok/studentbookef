package book_io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const bookDepartmentURL = api.BASE_URL + "bookdepartment/"

func CreateBookdepartment(bookdepartment domain.BookDepartment) (domain.BookDepartment, error) {
	entity := domain.BookDepartment{}
	resp, _ := api.Rest().SetBody(bookdepartment).Post(bookDepartmentURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookDepartment(id string) (domain.BookDepartment, error) {
	entity := domain.BookDepartment{}
	resp, _ := api.Rest().Get(bookDepartmentURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookDepartments() ([]domain.BookDepartment, error) {
	entity := []domain.BookDepartment{}
	resp, _ := api.Rest().Get(bookDepartmentURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookDepartment(bookdepartment domain.BookDepartment) (domain.BookDepartment, error) {
	entity := domain.BookDepartment{}
	resp, _ := api.Rest().SetBody(bookdepartment).Post(bookDepartmentURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBookDepartment(bookdepartment domain.BookDepartment) (domain.BookDepartment, error) {
	entity := domain.BookDepartment{}
	resp, _ := api.Rest().SetBody(bookdepartment).Post(bookDepartmentURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllOfBookDepartment(id string) ([]domain.BookDepartment, error) {
	entity := []domain.BookDepartment{}
	resp, _ := api.Rest().Get(bookDepartmentURL + "readAllOf?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
