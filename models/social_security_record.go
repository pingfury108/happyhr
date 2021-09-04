package models

import (
	"time"

	"gorm.io/gorm"
)

//社保缴纳记录
type SocialSecurityRecord struct {
	gorm.Model
	SerialNumber      uint
	StartPaymentTime  time.Time
	PensionBase       int
	MedicalBase       int
	WorkInJuryBase    int
	UnemploymentBase  int
	ProvidentFundBase int
}
