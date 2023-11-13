package utils

import (
	"bytes"
	"os"
	"html/template"

	"github.com/resendlabs/resend-go"
)

type EmailData struct {
	Title       string
	ContentData interface{}
	EmailTo     string
	Template    string
}

func SendEmail(data EmailData) error {
	
	var err error
	template, err := template.ParseFiles("utils/html_templates/" + data.Template)
	if err != nil {
		return err
	}
	var buf bytes.Buffer
	err = template.Execute(&buf, data.ContentData)
	if err != nil {
		return err
	}
	client := resend.NewClient(os.Getenv("EMAIL_API_KEY"))
	params := &resend.SendEmailRequest{
		From:    "Kura <team@kuragames.com>",
		To:      []string{data.EmailTo},
		Html:   buf.String(),
		Subject: "Welcome To SpeedQuest",
	}
	sent, err := client.Emails.Send(params)
	if err != nil {
		panic(err)
	}
	println(sent.Id)
	return err
}
