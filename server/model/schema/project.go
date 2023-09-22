package schema

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name string `gorm:"column:name"`
}
