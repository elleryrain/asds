package mail

import (
	"html/template"
	"strings"

	"git.ubrato.ru/ubrato/dispatch-service/internal/broker"
	"git.ubrato.ru/ubrato/dispatch-service/internal/pb/models/v1"
	"google.golang.org/protobuf/proto"
)

func (s *Service) resetPassHandler(msg *broker.Message) error {
	passwordRecovery := &models.PasswordRecovery{}
	err := proto.Unmarshal(msg.Data, passwordRecovery)
	if err != nil {
		return err
	}
	template, err := template.ParseFiles("./templates/resetPassword.templ")
	if err != nil {
		return err
	}
	body := new(strings.Builder)
	template.Execute(body, passwordRecovery)
	return s.smtpClient.Send(passwordRecovery.Email, "Сброс пароля", []byte(body.String()))
}
