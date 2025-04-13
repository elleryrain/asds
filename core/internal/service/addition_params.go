package service

type AdditionCreateParams struct {
	TenderID    int
	Title       string
	Content     string
	Attachments []string
}

type GetAdditionParams struct {
	TenderID int
}
