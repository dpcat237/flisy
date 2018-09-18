package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model

	Email     string
	Password  string
	Temporary bool
}

func (User) TableName() string {
	return "user"
}
