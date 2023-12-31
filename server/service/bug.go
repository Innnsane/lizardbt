package service

import (
	"lizardbt/global"
	"lizardbt/model/schema"
)

func CreateBug(bug *schema.Bug) (ret *schema.Bug, err error) {
	err = global.DB.Model(&schema.Bug{}).Create(bug).Error
	return
}

func QueryBug(bugId uint) (ret *schema.Bug, err error) {
	err = global.DB.Model(&schema.Bug{}).Where("id = ?", bugId).Find(&ret).Error
	return
}

func DeleteBug(bugId uint) (err error) {
	err = global.DB.Delete(&schema.Bug{}, "id = ?", bugId).Error
	return
}

func QueryBugs() (ret []*schema.Bug, err error) {
	err = global.DB.Model(&schema.Bug{}).Find(&ret).Error
	return
}
