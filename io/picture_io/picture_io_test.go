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
	result, err := ReadPicture("PF-f9a937ac-3a21-441f-af59-0ea39f7579fd")
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestReadPictures(t *testing.T) {
	result, err := ReadPictures()
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestUpdatePicture(t *testing.T) {

}
