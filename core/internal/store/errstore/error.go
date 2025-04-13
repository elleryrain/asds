package errstore

import "errors"

const (
	// foreign key violation: 23503
	FKViolation = "23503"
	// unique violation: 23505
	UniqueConstraint = "23505"
)

var (
	ErrQuestionnaireExist    = errors.New("questionnaire has been completed")
	ErrQuestionnaireNotFound = errors.New("questionnaire not found")
)

var (
	ErrUserNotFound = errors.New("user not found")
)

var (
	ErrTenderNotFound = errors.New("tender not found")
	ErrTenderAdditionNotFound = errors.New("tender addition not found")
)

var (
	ErrOrganizationNotFound       = errors.New("organization not found")
	ErrOrganizationNotAContractor = errors.New("organization is not a contractor")
)

var (
	ErrQuestionAnswerUniqueViolation = errors.New("answer to question already exists")
	ErrQuestionAnswerFKViolation     = errors.New("foreign key violation on question_answer")
	ErrQuestionAnswerNotFound        = errors.New("question-answer not found")
)

var (
	ErrPortfolioNotFound = errors.New("portfolio not found")
)

var (
	ErrFavouriteNotFound = errors.New("favourite not found")
)
