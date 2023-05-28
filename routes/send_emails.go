package routes

import (
	"crypto-api/helpers"
	"fmt"
	"net/smtp"
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendEmails(c *gin.Context)  {
	from := "gen.testtask@gmail.com"
	password := "nwvawyltynxwgviv"
	host := "smtp.gmail.com"
	port := "587"

	auth := smtp.PlainAuth("", from, password, host)
	emails, _ := helpers.ExtractEmailsFromFile("emails");
	subject := "Blockchain forever"
	body := "Price "+GetCurrentPrice()

	message := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", strings.Join(emails, ","), subject, body)

	smtp.SendMail(host+":"+port, auth, from, emails, []byte(message))
	c.JSON(http.StatusOK, []byte(""))
}