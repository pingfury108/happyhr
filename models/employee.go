package models

import (
	"time"

	"gorm.io/gorm"
)

// 员工
type Employee struct {
	gorm.Model
	SerialNumber          uint `grom:"unique"`
	Name                  string
	Gender                int
	Phone                 string
	DateBirth             time.Time
	Ethnic                string
	Hometown              string
	IdentityNumber        string `grom:"size:18;unique"`
	SalaryCardNumber      string
	MaritalStatus         bool
	HomeAddress           string
	PostCode              string
	HukuLocation          string
	HukuType              string
	Email                 string
	GraduatedSchool       string
	GraduationTime        time.Time
	Specialty             string
	EducationalBackground string
	ToWorkTime            time.Time
	OnboardingTime        time.Time
	TurnPositiveTime      time.Time // 转正时间
	EmergencyContact      string
	EmergencyContactPhone string
}
