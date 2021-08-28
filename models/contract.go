package models

import (
	"time"

	"gorm.io/gorm"
)

type Contract struct {
	gorm.Model
	SigningTime time.Time
	ExpireDate  time.Time
	Type        string
	Years       int
	Salary      int
	Job         string
}
