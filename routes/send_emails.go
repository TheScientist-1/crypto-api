package routes

import (
	"fmt"
	"net/smtp"
	"strings"
)

// Функція для відправки електронних листів
func SendEmails(emails []string, subject string, body string) error {
	from := "gen.testtask@example.com"
	password := "testtask"
	host := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("", from, password, host)

	message := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", strings.Join(emails, ","), subject, body)

	err := smtp.SendMail(host+":"+port, auth, from, emails, []byte(message))
	if err != nil {
		return err
	}

	return nil
}