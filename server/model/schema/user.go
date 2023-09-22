package schema

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"column:nickname"`
	Name     string `gorm:"column:name"`
	Role     int    `gorm:"column:role"`
}
