package models

import (
	"time"

	"gorm.io/gorm"
)

// 奖惩记录
type RewardPunishmentRecord struct {
	gorm.Model
	Name     string
	Date     time.Time
	Describe string
	Remark   string
}
