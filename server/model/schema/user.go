package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string
	Name     string
	Role     int
}
