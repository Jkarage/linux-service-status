package linuxservicestatus

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

func SendMail(service string) error {
	smtpServer := os.Getenv("SMTP_SERVER")
	smtpEmail := os.Getenv("SMTP_SERVER_EMAIL")
	mailPassword := os.Getenv("SMTP_PASSWORD")
	smtpPort, _ := strconv.Atoi(os.Getenv("SMTP_PORT"))

	m := gomail.NewMessage()
	m.SetHeaders(map[string][]string{
		"From": {"pepkarage@gmail.com"},
		"To":   {"josephbkarage@gmail.com"},
		// "Cc":      {"kalebu@neurotech.africa"},
		// "Bcc":     {"mkumboelia@gmail.com"},
		"Subject": {"Service Inactive"},
	})
	m.SetBody("text/plain", fmt.Sprintf("Service %s is not currently in its active status, Current time: %v", service, time.Now()))

	dialer := gomail.NewDialer(smtpServer, smtpPort, smtpEmail, mailPassword)
	if err := dialer.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
