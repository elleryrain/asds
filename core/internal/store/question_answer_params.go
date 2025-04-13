package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type CreateQuestionAnswerParams struct {
	TenderID             int
	AuthorOrganizationID int
	ParentID             models.Optional[int]
	Type                 models.QuestionAnswerType
	Content              string
}

type QuestionAnswerGetParams struct {
	QuestionAnswerIDs models.Optional[[]int]
}

// Получение вопросов-ответов с фильтром по ролям
type QuestionAnswerGetWithAccessParams struct {
	TenderID        int
	VerifedOnly     bool
	IsTenderCreator bool
	OrganizationID  models.Optional[int]
}

type QuestionAnswerVerifStatusUpdateParams struct {
	QuestionAnswerID   int
	VerificationStatus models.VerificationStatus
}
