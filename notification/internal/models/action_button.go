package models

import (
	"net/url"

	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
)

type ActionButton struct {
	Text   string `json:"text,omitempty"`
	Url    string `json:"url,omitempty"`
	Styled bool   `json:"styled,omitempty"`
}

func NewOptActionButton(ab *ActionButton) Optional[ActionButton] {
	if ab == nil {
		return Optional[ActionButton]{}
	}

	return NewOptional(*ab)
}

func stringToUrl(s string) url.URL {
	url, _ := url.Parse(s)
	return *url
}

func ConvertActionButtonToOptAPI(actionButton *ActionButton) api.OptActionButton {
	if actionButton == nil {
		return api.OptActionButton{}
	}

	return api.OptActionButton{
		Value: api.ActionButton{
			Text:   actionButton.Text,
			URL:    stringToUrl(actionButton.Url),
			Styled: actionButton.Styled,
		},
		Set: true,
	}
}
