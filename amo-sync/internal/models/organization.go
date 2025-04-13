package models

import (
	"time"
)

type Contact struct {
	Contact string
	Info    string
}

type Organization struct {
	ID         int
	BrandName  string
	FullName   string
	ShortName  string
	INN        string
	OKPO       string
	OGRN       string
	KPP        string
	TaxCode    string
	Address    string
	AvatarURL  string
	Emails     []Contact
	Phones     []Contact
	Messengers []Contact
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
