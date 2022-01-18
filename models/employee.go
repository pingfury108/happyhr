package models

import (
	"time"

	"gorm.io/gorm"
)

// 员工
type Employee struct {
	gorm.Model
	SerialNumber          uint      `grom:"unique"` // 编号
	Name                  string    // 姓名
	Gender                int       // 性别 女: 1, 男 2
	Phone                 string    // 电话
	DateBirth             time.Time `json:"date_birth"` // 出生日期
	Ethnic                string    // 民族
	Hometown              string    // 籍贯
	IdentityNumber        string    `grom:"size:18;unique"` // 身份证号码
	SalaryCardNumber      string    // 工资卡号
	MaritalStatus         bool      // 婚姻状况
	HomeAddress           string    // 家庭住址
	PostCode              string    // 邮编
	HukuLocation          string    // 户口所在地
	HukuType              string    // 户口性质
	Email                 string    // 电子邮箱
	GraduatedSchool       string    // 毕业院校
	GraduationTime        time.Time // 毕业时间
	Specialty             string    // 专业
	EducationalBackground string    // 学历
	ToWorkTime            time.Time // 参加工作时间
	OnboardingTime        time.Time // 入职时间
	TurnPositiveTime      time.Time // 转正时间
	EmergencyContact      string    // 紧急联系人
	EmergencyContactPhone string    // 紧急联系人电话
}
