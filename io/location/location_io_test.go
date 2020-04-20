package location

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentbookef/domain"
	"testing"
)

func TestCreateLocation(t *testing.T) {
	location := domain.Location{"", "cape town", "10.24488", "23.94857", "voila"}
	result, err := CreateLocation(location)
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestDeleteLocation(t *testing.T) {
	location := domain.Location{"", "cape town", "10.24488", "23.94857", "voila"}
	result, err := DeleteLocation(location)
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestReadLocation(t *testing.T) {
	result, err := ReadLocation("location")
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestReadLocations(t *testing.T) {
	result, err := ReadLocations()
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
func TestUpdateLocation(t *testing.T) {
	location := domain.Location{"", "cape town", "10.24488", "23.94857", "voila"}
	result, err := UpdateLocation(location)
	assert.Nil(t, err)
	fmt.Println("result :", result)
}
