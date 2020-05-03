package user

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const userpostURL = api.BASE_URL + "user_post/"

func CreateUserPost(post domain.UserPost) (domain.UserPost, error) {
	entity := domain.UserPost{}
	resp, _ := api.Rest().SetBody(post).Post(userpostURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteUserPost(post domain.UserPost) (domain.UserPost, error) {
	entity := domain.UserPost{}
	resp, _ := api.Rest().SetBody(post).Post(userpostURL + "delete")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateUserPost(post domain.UserPost) (domain.UserPost, error) {
	entity := domain.UserPost{}
	resp, _ := api.Rest().SetBody(post).Post(userpostURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadUserPost(id string) (domain.UserPost, error) {
	entity := domain.UserPost{}
	resp, _ := api.Rest().Get(userpostURL + "read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUserPosts() ([]domain.UserPost, error) {
	entity := []domain.UserPost{}
	resp, _ := api.Rest().Get(userpostURL + "reads")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadAllOfUserPost(id string) ([]domain.UserPost, error) {
	entity := []domain.UserPost{}
	resp, _ := api.Rest().Get(userpostURL + "readAllOf?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)

	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
