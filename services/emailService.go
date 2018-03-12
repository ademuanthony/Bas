package services

import (
	"gopkg.in/gomail.v2"
	"crypto/tls"
	"github.com/ademuanthony/bas/resources"
	"fmt"
)

type EmailService struct {

}


func (this EmailService) SendEmail(from, to, title, body, contentType string) error {
	if contentType == "" {
		contentType = "text/html"
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody(contentType, body)


	d := gomail.NewDialer("smtp.superfluxnigeria.com", 587,
		"helpdesk@superfluxnigeria.com",
		"Software_2017")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		//panic(err)
		fmt.Println("Email not sent to " + to)
		fmt.Printf("Error: %v\n", err)
		return err
	}
	fmt.Println("Email sent to " + to)
	return nil
}

func (this EmailService) SendMultipleEmails(emails []resources.EmailResource) {
	for _, email := range emails{
		this.SendEmail(email.From, email.To, email.Title, email.Message, email.ContentType)
	}
}

type Email struct {
	From string 		`json:"from"`
	To string 		`json:"to"`
	Title string		`json:"title"`
	Body string		`json:"body"`
	ContentType string	`json:"content_type"`
}

type EmailUser struct {
	Username    string
	Password    string
	EmailServer string
	Port        int
}
