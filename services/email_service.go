package services

import (
	"net/smtp"

	"github.com/asmusj224/exercise-automation/config"
)

type EmailService interface {
	SendEmail(toEmailAddress, subject, body string) (string, error)
}

type emailService struct{}

func NewEmailService() EmailService {
	return &emailService{}
}

func (*emailService) SendEmail(toEmailAddress, subject, body string) (string, error) {
	from := config.Config("EMAIL_ADDRESS")
	password := config.Config("EMAIL_PASSWORD")
	to := []string{toEmailAddress}
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	msg := "From: " + from + "\n" +
		"To: " + toEmailAddress + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	message := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return "", err
	}

	return "Message sent", nil

}
