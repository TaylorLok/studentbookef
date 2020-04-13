package user

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const useraccountURL = api.BASE_URL + "user"

func UserLog(loginDetails domain.UserAccount) (domain.UserAccount, error) {
	var entity domain.UserAccount
	resp, _ := api.Rest().SetBody(loginDetails).Post(useraccountURL + "login")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity)
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}
