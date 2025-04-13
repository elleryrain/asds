package handler

import (
	modelsv1 "gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb/models/v1"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
)

func ConvertProtoUserToAmoModels(protoUser *modelsv1.User) (models.AmoContact, models.AmoCompany) {
	amoUser := models.AmoContact{
		FirstName:  protoUser.FirstName,
		LastName:   protoUser.LastName,
		MiddleName: protoUser.MiddleName,
		Phone:      protoUser.Phone,
		Email:      protoUser.Email,
		PhoneType:  models.AmoContactPhoneTypeWork,
		EmailType:  models.AmoContactEmailTypeWork,
	}

	var role models.AmoCompanyRole

	if protoUser.GetOrganization().GetIsContractor() {
		role = models.AmoCompanyRoleContractor
	} else {
		role = models.AmoCompanyRoleClient
	}

	amoCompany := models.AmoCompany{
		BrandName:    protoUser.Organization.BrandName,
		FullName:     protoUser.Organization.FullName,
		ShortName:    protoUser.Organization.ShortName,
		INN:          protoUser.Organization.Inn,
		OKPO:         protoUser.Organization.Okpo,
		OGRN:         protoUser.Organization.Ogrn,
		KPP:          protoUser.Organization.Kpp,
		TaxCode:      protoUser.Organization.TaxCode,
		Address:      protoUser.Organization.Address,
		Role:         role,
		Phone:        protoUser.Phone,
		Email:        protoUser.Email,
		RegisteredAt: protoUser.Organization.CreatedAt.AsTime(),
		Phone1:       getContact(protoUser.Organization.Phones, 0),
		Phone2:       getContact(protoUser.Organization.Phones, 1),
		Phone3:       getContact(protoUser.Organization.Phones, 2),
		Phone4:       getContact(protoUser.Organization.Phones, 3),
		Phone5:       getContact(protoUser.Organization.Phones, 4),
		Phone6:       getContact(protoUser.Organization.Phones, 5),
		Phone7:       getContact(protoUser.Organization.Phones, 6),
		Phone8:       getContact(protoUser.Organization.Phones, 7),
		Phone9:       getContact(protoUser.Organization.Phones, 8),
	}

	return amoUser, amoCompany
}

func getContact(contacts []*modelsv1.Contact, index int) string {
	if index < len(contacts) {
		return contacts[index].Contact
	}
	return ""
}
