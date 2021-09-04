package models

import "gorm.io/gorm"

//员工信息附件
type EmployeeInformationAttachment struct {
	gorm.Model
	IDCard                []byte
	DegreeCertificate     []byte
	GraduationCertificate []byte
	SalaryCard            []byte
	ElectronicaResume     []byte
	PaperFileFiling       []byte
}
