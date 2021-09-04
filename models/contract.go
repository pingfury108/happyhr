package models

import (
	"time"

	"gorm.io/gorm"
)

//合同
type Contract struct {
	gorm.Model
	SigningTime time.Time
	ExpireDate  time.Time
	Type        string
	Years       int
	Salary      int
	Job         string
}
