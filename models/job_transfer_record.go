package models

import (
	"time"

	"gorm.io/gorm"
)

type JobTransferRecord struct {
	gorm.Model
	Date               time.Time
	OriginalDepartment string
	OriginalJob        string
	ToDepartment       string
	ToJob              string
	Reason             string
}
