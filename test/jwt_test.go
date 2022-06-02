package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"simpledy/model"
	"simpledy/utils"
	"testing"
)

//var user = model.User{
//	Name:     "zhangsan",
//	Password: "12345678",
//}
var user = model.User{
	Model: gorm.Model{ID: 1},
	Name:  "zhangsan",
}

func TestCreateToken(t *testing.T) {
	tokenString, err := utils.CreateToken(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(tokenString)
}

func TestParseToken(t *testing.T) {
	tokenString, err := utils.CreateToken(user)
	if err != nil {
		panic(err)
	}
	_, claims := utils.ParseToken(tokenString)

	assert.Equal(t, claims["id"], float64(1))
	assert.Equal(t, claims["username"], "zhangsan")
}