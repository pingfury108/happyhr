package models

import "gorm.io/gorm"

//个人简历
type Resume struct {
	gorm.Model
	WorkExperience        string
	EducationalExperience string
}
