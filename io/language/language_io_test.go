package language

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentbookef/domain"
	"testing"
)

func TestCreateLanguage(t *testing.T) {
	language := domain.Language{"", "French"}
	result, err := CreateLanguage(language)
	assert.Nil(t, err)
	fmt.Println(result)
}


func TestReadLanguage(t *testing.T) {
	result, err := ReadLanguage("")
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestDeleteLanguage(t *testing.T) {
	language := domain.Language{"", "french"}
	result, err := DeleteLanguage(language)
	assert.Nil(t, err)
	fmt.Println(result)
}
func TestReadLanguages(t *testing.T) {
	result, err := ReadLanguages()
	assert.Nil(t, err)
	fmt.Println(result)
}
