package book_io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const bookImageURL = api.BASE_URL + "bookimage/"

func CreatBookImage(bookImage domain.BookImage) (domain.BookImage, error) {
	entity := domain.BookImage{}
	resp, _ := api.Rest().SetBody(bookImage).Post(bookImageURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookImage(id string) (domain.BookImage, error) {
	entity := domain.BookImage{}
	resp, _ := api.Rest().Get(bookImageURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookImages() ([]domain.BookImage, error) {
	entity := []domain.BookImage{}
	resp, _ := api.Rest().Get(bookImageURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateBookImage(bookImage domain.BookImage) (domain.BookImage, error) {
	entity := domain.BookImage{}
	resp, _ := api.Rest().SetBody(bookImage).Post(bookImageURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookImage(bookImage domain.BookImage) (domain.BookImage, error) {
	entity := domain.BookImage{}
	resp, _ := api.Rest().SetBody(bookImage).Post(bookImageURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
