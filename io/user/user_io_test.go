package user

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"studentbookef/domain"
	"testing"
	"time"
)

func TestUserLog(t *testing.T) {
	userAccount := domain.UserAccount{"espoir@admin.com", "HQfeJib6", "", time.Now()}
	result, err := UserLog(userAccount)

	fmt.Println(result)
	assert.Nil(t, err)
}
func TestReadAllLog(t *testing.T) {
	result, err := ReadAllLog()
	fmt.Println(result)
	assert.Nil(t, err)
}
