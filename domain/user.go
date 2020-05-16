package domain

import "time"

type User struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	PhoneNumber string `json:"phoneNumber"`
}
type UserDepartment struct {
	Email        string `json:"email"`
	DepartmentId string `json:"department_id"`
	Description  string `json:"description"`
}
type UserImage struct {
	Email       string    `json:"email"`
	ImageId     string    `json:"image_id"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}
type UserAccount struct {
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	AccountStatus string    `json:"account_status"`
	Date          time.Time `json:"date"`
}
type UserPost struct {
	PostId string `json:"postId"`
	Email  string `json:"email"`
}
type UserProfile struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	PhoneNumber     string `json:"phoneNumber"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm-password"`
}
