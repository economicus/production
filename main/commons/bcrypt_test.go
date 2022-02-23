package commons

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	pwd := "123456"
	a := assert.New(t)
	_, err := HashPassword([]byte(pwd))
	a.NoError(err, fmt.Sprintf("error must not be generated: %v", err))
}

func TestComparePassword(t *testing.T) {
	pwd := "123456"
	a := assert.New(t)
	hashedPwd, _ := HashPassword([]byte(pwd))
	validPwd := "123456"
	invalidPwd := "12345"
	err := ComparePassword([]byte(validPwd), hashedPwd)
	a.NoError(err, fmt.Sprintf("error must not be generated: %v", err))
	err = ComparePassword([]byte(invalidPwd), hashedPwd)
	a.Error(err, fmt.Sprintf("error must be generated: %v", err))
}
