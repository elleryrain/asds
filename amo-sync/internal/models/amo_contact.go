package models

type AmoContactPhoneType int

const (
	AmoContactPhoneTypeWork   AmoContactPhoneType = 338983
	AmoContactPhoneTypeWorkDD AmoContactPhoneType = 338985
	AmoContactPhoneTypeMob    AmoContactPhoneType = 338987
	AmoContactPhoneTypeFax    AmoContactPhoneType = 338989
	AmoContactPhoneTypeHome   AmoContactPhoneType = 338991
	AmoContactPhoneTypeOther  AmoContactPhoneType = 338993
)

type AmoContactEmailType int

const (
	AmoContactEmailTypeWork  AmoContactEmailType = 338995
	AmoContactEmailTypePriv  AmoContactEmailType = 338997
	AmoContactEmailTypeOther AmoContactEmailType = 338999
)

type AmoContact struct {
	// ID         int64
	FirstName  string
	LastName   string
	MiddleName string
	Position   string
	Phone      string
	Email      string
	Telegram   string
	WA         string
	PhoneType  AmoContactPhoneType
	EmailType  AmoContactEmailType
}
