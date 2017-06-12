package services

import (
	"gopkg.in/gomail.v2"
	"crypto/tls"
)

type EmailService struct {

}


func (this EmailService) SendEmail(from, to, title, body, contentType string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody(contentType, body)


	d := gomail.NewDialer("smtp.superfluxnigeria.com", 587,
		"anthony_ademu@superfluxnigeria.com",
		"Ademu_17")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		//panic(err)
		return err
	}
	return nil
}


type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}
