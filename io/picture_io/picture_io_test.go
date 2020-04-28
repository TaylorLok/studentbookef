package picture_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentbookef/domain"
	"testing"
)

func TestCreatePicture(t *testing.T) {
	picture := domain.Picture{}
	result, err := CreatePicture(picture)
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestReadPicture(t *testing.T) {
	result, err := ReadPicture("PF-1ec1d101-16c0-4dfe-ada6-b160a3e5a073")
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestReadPictures(t *testing.T) {
	result, err := ReadPictures()
	assert.Nil(t, err)
	//fmt.Println("result :", result)
	for _, picture := range result {
		fmt.Println(picture.Id)
	}
}
func TestUpdatePicture(t *testing.T) {

}
func TestReadFirstPicture(t *testing.T) {
	result, err := ReadFirstPicture("PF-85a05436-907d-451f-8cbd-ec17c5332038")
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestReadAllOf(t *testing.T) {

	ids := []string{"PF-00bfba1a-1b23-43f6-be1b-e2a177025006", "PF-1ec1d101-16c0-4dfe-ada6-b160a3e5a073"}
	result, err := ReadAllOf(ids)
	assert.Nil(t, err)
	//fmt.Println("result :", result)
	for _, picture := range result {
		fmt.Println(picture.Id, "   ", picture.Description)
	}
}
