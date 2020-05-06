package io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const locationURL = api.BASE_URL + "location/"

func CreateLocation(location domain.Location) (domain.Location, error) {
	entity := domain.Location{}
	resp, _ := api.Rest().SetBody(location).Post(locationURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadLocation(id string) (domain.Location, error) {
	entity := domain.Location{}
	resp, _ := api.Rest().Get(locationURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func UpdateLocation(location domain.Location) (domain.Location, error) {
	entity := domain.Location{}
	resp, _ := api.Rest().SetBody(location).Post(locationURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func DeleteLocation(location domain.Location) (domain.Location, error) {
	entity := domain.Location{}
	resp, _ := api.Rest().SetBody(location).Post(locationURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
