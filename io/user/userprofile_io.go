package user

import (
	"errors"
	"studentbookef/api"
	"studentbookef/domain"
)

const UserprofileURL = api.BASE_URL + "userprofile/"

func CreateUserProfile(user domain.User) (domain.User, error) {
	var entity domain.User
	resp, _ := api.Rest().SetBody(user).Post(UserprofileURL + "create")
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity) //there is an error in this line "Unmarshal" that what the errors saying
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

func ReadUserProfile(email string) (domain.User, error) {
	var entity domain.User
	resp, _ := api.Rest().Get(UserprofileURL + "read?id=" + email)
	if resp.IsError() {
		return entity, errors.New(resp.Status())
	}
	err := api.JSON.Unmarshal(resp.Body(), &entity) //there is an error in this line "Unmarshal" that what the errors saying
	if err != nil {
		return entity, errors.New(resp.Status())
	}
	return entity, nil
}

//func UpdateUserProfile(User domain.UserProfile) (domain., error) { //// I wonder why i can't pass my domain here? domain.Userprofile its giving errors
//	entity := domain.UserProfile{}
//	resp, _ := api.Rest().SetBody(post).Post(UserprofileURL + "update")
//
//	if resp.IsError() {
//		return entity, errors.New(resp.Status())
//	}
//	err := api.JSON.Unmarshal(resp.Body(), &entity)
//
//	if err != nil {
//		return entity, errors.New(resp.Status())
//	}
//	return entity, nil
//}
