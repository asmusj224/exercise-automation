package services

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/asmusj224/exercise-automation/config"
)

type MessageService interface {
	SendSMS(toPhoneNumber, message string) (string, error)
}

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	client HttpClient
)

type messageService struct{}

type Message struct {
	Sid string `json:"sid"`
}

func NewMessageService(c HttpClient) MessageService {
	client = c
	return &messageService{}
}

func (*messageService) SendSMS(toPhoneNumber, message string) (string, error) {
	accountSid := config.Config("ACCOUNT_SID")
	authToken := config.Config("AUTH_TOKEN")
	fromPhoneNumber := config.Config("FROM_PHONE_NUMBER")
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	msgData := url.Values{}
	msgData.Set("To", toPhoneNumber)
	msgData.Set("From", fromPhoneNumber)
	msgData.Set("Body", message)
	msgDataReader := *strings.NewReader(msgData.Encode())

	req, err := http.NewRequest("POST", urlStr, &msgDataReader)
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	var target Message
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&target)
	return target.Sid, nil
}
