package schema

import "gorm.io/gorm"

type Bug struct {
	gorm.Model
	ReporterId      uint    `gorm:"column:reporter_id"`
	Reporter        User    `gorm:"foreignKey:ReporterID;references:ID"`
	HandlerId       uint    `gorm:"column:handler_id"`
	Handler         User    `gorm:"foreignKey:HandlerID;references:ID"`
	Status          int     `gorm:"column:status"`
	Priority        int     `gorm:"column:priority"`
	Severity        int     `gorm:"column:severity"`
	Reproducibility int     `gorm:"column:reproducibility"`
	ProjectId       uint    `gorm:"column:project_id"`
	Project         Project `gorm:"foreignKey:ProjectID;references:ID"`
	VersionId       uint    `gorm:"column:version_id"`
	Version         Version `gorm:"foreignKey:VersionID;references:ID"`
	GroupId         uint    `gorm:"column:group_id"`
	Group           Group   `gorm:"foreignKey:GroupID;references:ID"`
	Summary         string  `gorm:"column:summary"`
	Description     string  `gorm:"column:description"`
}
