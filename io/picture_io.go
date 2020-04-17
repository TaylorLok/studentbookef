package io

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const pictureURL = api.BASE_URL +"picture/"

func CreatePicture(picture domain.Picture)(domain.Picture,error)  {
	entity:= domain.Picture{}
	resp,_ := api.Rest().SetBody(picture).Post(pictureURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadPicture(id string)(domain.Picture,error)  {
	entity := domain.Picture{}
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

func UpdatePicture(picture domain.Picture)(domain.Picture,error)  {
	entity := domain.Picture{}
	resp, _ := api.Rest().SetBody(picture).Post(pictureURL + "update")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func DeletePicture(picture domain.Picture)(domain.Picture,error){
	entity := domain.Picture{}
	resp, _ := api.Rest().SetBody(picture).Post(pictureURL + "delete")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
