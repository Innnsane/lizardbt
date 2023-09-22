package schema

import "gorm.io/gorm"

type Bug struct {
	gorm.Model
	ReporterID      uint
	Reporter        User
	HandlerID       uint
	Handler         User
	Status          int
	Priority        int
	Severity        int
	Reproducibility int
	VersionID       uint
	Version         Version
	GroupID         uint
	Group           Group
	Summary         string
	Description     string
}
