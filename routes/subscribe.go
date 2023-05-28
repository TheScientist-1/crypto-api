package routes

import (

  "net/http"
  "net/mail"
  "crypto-api/helpers"
  "github.com/gin-gonic/gin"
)

type Email struct {
  Email string `json:"email"`
}

func Subscribe(c *gin.Context) {
  var email Email

  c.BindJSON(&email)
  if _, err := mail.ParseAddress(email.Email); err == nil {
    helpers.SaveEmail(email.Email)
  } else {
    c.IndentedJSON(http.StatusConflict, err.Error())
  }

  c.IndentedJSON(http.StatusOK, "")
}