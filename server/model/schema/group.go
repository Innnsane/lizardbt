package schema

import "gorm.io/gorm"

type Group struct {
	gorm.Model
	Name string `gorm:"column:name"`
}
