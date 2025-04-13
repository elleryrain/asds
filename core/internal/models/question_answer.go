package models

import (
	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

type QuestionAnswerType int

const (
	QuestionAnswerTypeInvalid QuestionAnswerType = iota
	QuestionAnswerTypeQuestion
	QuestionAnswerTypeAnswer
)

var mapQuestionAnswerType = map[QuestionAnswerType]api.QuestionAnswerType{
	QuestionAnswerTypeQuestion: api.QuestionAnswerTypeQuestion,
	QuestionAnswerTypeAnswer:   api.QuestionAnswerTypeAnswer,
}

func (q QuestionAnswerType) ToAPI() api.QuestionAnswerType {
	if apiType, ok := mapQuestionAnswerType[q]; ok {
		return apiType
	}
	return api.QuestionAnswerTypeInvalid
}

var mapApiToQuestionAnswerType = map[api.QuestionAnswerType]QuestionAnswerType{
	api.QuestionAnswerTypeQuestion: QuestionAnswerTypeQuestion,
	api.QuestionAnswerTypeAnswer:   QuestionAnswerTypeAnswer,
}

func APIToQuestionAnswerType(apiType api.QuestionAnswerType) QuestionAnswerType {
	questionAnswerType, ok := mapApiToQuestionAnswerType[apiType]
	if !ok {
		return QuestionAnswerTypeInvalid
	}
	return questionAnswerType
}

type QuestionAnswer struct {
	ID                   int
	TenderID             int
	AuthorOrganizationID int
	ParentID             Optional[int]
	Type                 QuestionAnswerType
	Content              string
	VerificationStatus   VerificationStatus
}

func ConvertQuestionAnswerToAPI(qe QuestionAnswer) api.QuestionAnswer {
	return api.QuestionAnswer{
		ID:                 qe.ID,
		ParentID:           api.OptInt{Value: qe.ParentID.Value, Set: qe.ParentID.Set},
		Type:               qe.Type.ToAPI(),
		Content:            qe.Content,
		VerificationStatus: qe.VerificationStatus.ToAPI(),
	}
}

type QuestionWithAnswer struct {
	VerificationObject

	Question QuestionAnswer
	Answer   Optional[QuestionAnswer]
}

func (qe QuestionWithAnswer) ToVerificationObject() api.VerificationRequestObject {
	return api.VerificationRequestObject{
		Type:               api.QuestionWithAnswerVerificationRequestObject,
		QuestionWithAnswer: ConvertQuestionWithAnswerToAPI(qe),
	}
}

func ConvertQuestionWithAnswerToAPI(qe QuestionWithAnswer) api.QuestionWithAnswer {
	return api.QuestionWithAnswer{
		Question: ConvertQuestionAnswerToAPI(qe.Question),
		Answer:   api.OptQuestionAnswer{Value: ConvertQuestionAnswerToAPI(qe.Answer.Value), Set: qe.Answer.Set},
	}
}
