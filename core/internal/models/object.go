package models

import api "gitlab.ubrato.ru/ubrato/core/api/gen"

type ObjectType int

const (
	ObjectTypeInvalid ObjectType = iota
	ObjectTypeOrganization
	ObjectTypeTender
	ObjectTypeAddition
	ObjectTypeQuestionAnswer
)

var mapVerificationObjectType = map[ObjectType]api.ObjectType{
	ObjectTypeInvalid:        api.ObjectTypeInvalid,
	ObjectTypeOrganization:   api.ObjectTypeOrganization,
	ObjectTypeTender:         api.ObjectTypeTender,
	ObjectTypeAddition:       api.ObjectTypeAddition,
	ObjectTypeQuestionAnswer: api.ObjectTypeQuestionAnswer,
}

func (v ObjectType) ToAPI() api.ObjectType {
	return mapVerificationObjectType[v]
}
