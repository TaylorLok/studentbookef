package language

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const languageURL = api.BASE_URL + "language/"

func CreateLanguage(language domain.Language) (domain.Language, error) {
	entity := domain.Language{}
	resp, _ := api.Rest().SetBody(language).Post(languageURL + "create")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(err.Error())
	}
	return entity, nil
}
func DeleteLanguage(language domain.Language) (domain.Language, error) {
	entity := domain.Language{}
	resp, _ := api.Rest().SetBody(language).Post(languageURL + "delete")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(err.Error())
	}
	return entity, nil
}
func UpdateLanguage(language domain.Language) (domain.Language, error) {
	entity := domain.Language{}
	resp, _ := api.Rest().SetBody(language).Post(languageURL + "update")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(err.Error())
	}
	return entity, nil
}
func ReadLanguage(id string) (domain.Language, error) {
	entity := domain.Language{}
	resp, _ := api.Rest().Get(languageURL + "read?id=" + id)

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(err.Error())
	}
	return entity, nil
}
func ReadLanguages() ([]domain.Language, error) {
	entity := []domain.Language{}
	resp, _ := api.Rest().Get(languageURL + "reads")

	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(err.Error())
	}
	return entity, nil
}
