package schema

import "gorm.io/gorm"

type Version struct {
	gorm.Model
	Name string `gorm:"column:name"`
}
