package user

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Uuid string
}
