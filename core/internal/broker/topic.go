package broker

type Topic string

const (
	UbratoUserRegisteredSubject  = "ubrato.user.registered"
	UbratoSurveySubmittedSubject = "ubrato.survey.submitted"
	UbratoUserConfirmEmail       = "email.send.confirmation"
	UbratoUserEmailResetPass     = "email.send.resetpass"

	NotifyUserEmailConfirmation = "notify.user.email.confirmation"
	NotifyUserEmailConfirmed    = "notify.user.email.confirmed"

	NotifyOrganizationVerification         = "notify.organization.verification"
	NotifyTenderVerification               = "notify.tender.verification"
	NotifyTenderAdditionVerification       = "notify.tender.addition.verification"
	NotifyTenderQuestionAnswerVerification = "notify.tender.question.answer.verification"

	NotifyTenderQA      = "notify.tender.qa"
	NotifyTenderWinners = "notify.tender.winners"
)
