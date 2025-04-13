package models

import (
	"time"
)

type AmoCompanyRole int

const (
	AmoCompanyRoleClient     AmoCompanyRole = 374361
	AmoCompanyRoleContractor AmoCompanyRole = 374363
)

type AmoCompanyPhoneType int

const (
	AmoCompanyPhoneTypeWork   AmoCompanyPhoneType = 338983
	AmoCompanyPhoneTypeWorkDD AmoCompanyPhoneType = 338985
	AmoCompanyPhoneTypeMob    AmoCompanyPhoneType = 338987
	AmoCompanyPhoneTypeFax    AmoCompanyPhoneType = 338989
	AmoCompanyPhoneTypeHome   AmoCompanyPhoneType = 338991
	AmoCompanyPhoneTypeOther  AmoCompanyPhoneType = 338993
)

type AmoCompanyEmailType int

const (
	AmoCompanyEmailTypeWork  AmoCompanyEmailType = 338995
	AmoCompanyEmailTypePriv  AmoCompanyEmailType = 338997
	AmoCompanyEmailTypeOther AmoCompanyEmailType = 338999
)

type AmoCompany struct {
	// ID                 int64
	Phone              string
	Email              string
	Web                string
	Address            string
	Role               AmoCompanyRole
	FullName           string
	ShortName          string
	INN                string
	KPP                string
	RegisteredAt       time.Time
	Region             string
	OKVED2             string
	MainActivity       string
	OKPO               string
	OGRN               string
	TaxCode            string
	LegalAddress       string
	City               string
	WorkType           string
	ActivityField      string
	Comment            string
	CompanyDescription string
	BrandName          string
	Phone1             string
	Phone2             string
	Phone3             string
	Phone4             string
	Phone5             string
	Phone6             string
	Phone7             string
	Phone8             string
	Phone9             string
	PhoneType          AmoCompanyPhoneType
	EmailType          AmoCompanyEmailType
}
