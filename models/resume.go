package models

import "gorm.io/gorm"

type Resume struct {
	gorm.Model
	WorkExperience        string
	EducationalExperience string
}
