package book_io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const bookPostURL = api.BASE_URL + "bookpost/"

func CreatBookPost(post domain.BookPost) (domain.BookPost, error) {
	entity := domain.BookPost{}
	resp, _ := api.Rest().SetBody(post).Post(bookPostURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteBookPost(post domain.BookPost) (domain.BookPost, error) {
	entity := domain.BookPost{}
	resp, _ := api.Rest().SetBody(post).Post(bookPostURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateBookPost(post domain.BookPost) (domain.BookPost, error) {
	entity := domain.BookPost{}
	resp, _ := api.Rest().SetBody(post).Post(bookPostURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookPost(id string) (domain.BookPost, error) {
	entity := domain.BookPost{}
	resp, _ := api.Rest().Get(bookPostURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadWithBookId(id string) (domain.BookPost, error) {
	entity := domain.BookPost{}
	resp, _ := api.Rest().Get(bookPostURL + "readWithbookId?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadBookPosts() ([]domain.BookPost, error) {
	entity := []domain.BookPost{}
	resp, _ := api.Rest().Get(bookPostURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
