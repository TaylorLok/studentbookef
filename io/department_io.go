package io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const departmentURL = api.BASE_URL + "department/"

func CreateDepartment(department domain.Department) (domain.Department, error) {
	entity := domain.Department{}
	resp, _ := api.Rest().SetBody(department).Post(departmentURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadDepartment(id string) (domain.Department, error) {
	entity := domain.Department{}
	resp, _ := api.Rest().Get(departmentURL + "read?id=" + id)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
func ReadDepartments() ([]domain.Department, error) {
	entity := []domain.Department{}
	resp, _ := api.Rest().Get(departmentURL + "reads")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func UpdateDepartment(department domain.Department) (domain.Department, error) {
	entity := domain.Department{}
	resp, _ := api.Rest().SetBody(department).Post(departmentURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeleteDepartment(department domain.Department) (domain.Department, error) {
	entity := domain.Department{}
	resp, _ := api.Rest().SetBody(department).Post(departmentURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
