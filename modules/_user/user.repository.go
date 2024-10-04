package user

import (
	"go-boiler-plate/infra/database"
	"go-boiler-plate/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) (*User, error) {
	db := database.GetDb(c)

	user := User{}

	result := db.Find(&user, &User{Model: gorm.Model{ID: 8}})

	return util.HandleDbResponseError(result, "not found", &user)
}
