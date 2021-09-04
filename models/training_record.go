package models

import "gorm.io/gorm"

// 培训记录
type TrainingRecord struct {
	gorm.Model
	Name         string
	Address      string
	TrainingTime string
	Content      string
	Result       string
}
