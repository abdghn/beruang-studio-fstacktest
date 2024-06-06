package controllers

import (
	"backend/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"log"
	"os"
)

type AdminController struct {
	DB *gorm.DB
}

func (ctrl *AdminController) SendInvitation(c *gin.Context) {
	var input struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	link := generateInvitationLink()
	invitation := models.Invitation{Email: input.Email, Link: link, Status: "Pending"}
	ctrl.DB.Create(&invitation)

	// Send Email
	sendEmail(input.Email, link)

	c.JSON(200, gin.H{"message": "Invitation sent"})
}

func generateInvitationLink() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func sendEmail(to string, link string) {
	from := os.Getenv("EMAIL_USERNAME")
	msg := "Please register using the following link: " + link

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("Info Admin <%s>", from))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Invitation")
	mailer.SetBody("text/plain", msg)
	dialer := gomail.NewDialer(
		os.Getenv("EMAIL_HOST"),
		587,
		from,
		os.Getenv("EMAIL_PASSWORD"),
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Println("send email : ", err)
		return
	}

}
