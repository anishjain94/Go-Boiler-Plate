package user

import (
	"context"
	"go-boiler-plate/infra/database"
	"go-boiler-plate/util"

	"gorm.io/gorm"
)

func GetUser(ctx *context.Context) (*User, error) {

	db := database.GetDb(ctx)

	user := User{}

	result := db.Find(&user, &User{Model: gorm.Model{ID: 8}})

	return util.HandleDbResponseError(result, "not found", &user)
}
