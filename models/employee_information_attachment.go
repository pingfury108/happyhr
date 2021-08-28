package models

import "gorm.io/gorm"

type EmployeeInformationAttachment struct {
	gorm.Model
	IDCard                []byte
	DegreeCertificate     []byte
	GraduationCertificate []byte
	SalaryCard            []byte
	ElectronicaResume     []byte
	PaperFileFiling       []byte
}
