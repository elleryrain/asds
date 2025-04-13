package smtp

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/mail"
	"net/smtp"
)

type client struct {
	auth smtp.Auth

	emailFrom string
	host      string
}

type SMTPClient interface {
	Send(emailTo string, subject string, body []byte) (err error)
}

func NewSMTPClient(login string, password string, host string, from string) (*client, error) {
	hostx, _, _ := net.SplitHostPort(host)
	auth := smtp.PlainAuth("", login, password, hostx)

	return &client{
		auth:      auth,
		emailFrom: from,
		host:      host,
	}, nil
}

func (s *client) Send(emailTo string, subject string, body []byte) (err error) {
	from := mail.Address{Name: "noreply", Address: s.emailFrom}
	to := mail.Address{Name: "", Address: emailTo}

	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += string(body)

	host, _, _ := net.SplitHostPort(s.host)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(s.host)
	if err != nil {
		return err
	}

	c.StartTLS(tlsconfig)

	if err = c.Auth(s.auth); err != nil {
		return err
	}

	if err = c.Mail(from.Address); err != nil {
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		return err
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	c.Quit()

	return nil
}
