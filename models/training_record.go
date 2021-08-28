package models

import "gorm.io/gorm"

type TrainingRecord struct {
	gorm.Model
	Name         string
	Address      string
	TrainingTime string
	Content      string
	Result       string
}
