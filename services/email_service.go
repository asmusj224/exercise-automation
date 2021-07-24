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

	message := []byte(subject + body)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return "", err
	}

	return "Message sent", nil

}

// 	// message
// 	subject := "Subject: Our Golang Email\n"
// 	body := "our first email!"
// 	message := []byte(subject + body)
// 	// athentication data
// 	// func PlainAuth(identity, username, password, host string) Auth
// 	auth := smtp.PlainAuth("", from, password, host)
// 	// send mail
// 	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
// 	err := smtp.SendMail(address, auth, from, to, message)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}
// }
