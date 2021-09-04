package models

import (
	"time"

	"gorm.io/gorm"
)

//调岗记录
type JobTransferRecord struct {
	gorm.Model
	Date               time.Time
	OriginalDepartment string
	OriginalJob        string
	ToDepartment       string
	ToJob              string
	Reason             string
}
