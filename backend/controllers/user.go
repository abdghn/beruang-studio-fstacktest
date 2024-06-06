package controllers

import (
	"backend/models"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/skip2/go-qrcode"
	"gopkg.in/gomail.v2"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func (ctrl *UserController) Register(c *gin.Context) {
	var input struct {
		Email         string `json:"email"`
		Nama          string `json:"nama"`
		TanggalLahir  string `json:"tanggal_lahir"`
		JenisKelamin  string `json:"jenis_Kelamin"`
		AlamatLengkap string `json:"alamat_lengkap"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Generate registration code and QR code
	registrationCode := generateRegistrationCode()
	qrCode, err := generateQRCode(registrationCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate QR code"})
		return
	}

	// Create registration record
	registration := models.Registration{
		Email:            input.Email,
		RegistrationCode: registrationCode,
		QRCode:           qrCode,
		Status:           "Registered",
	}
	if err := ctrl.DB.Create(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create registration record"})
		return
	}

	go scheduleEmailVerification(input.Email)

	c.JSON(http.StatusOK, gin.H{
		"message":          "Registration successful",
		"registrationCode": registrationCode,
		"qrCode":           qrCode,
	})
}

func (ctrl *UserController) GetInvitation(c *gin.Context) {
	link := c.Param("link")
	var invitation models.Invitation
	if err := ctrl.DB.Where("link = ?", link).First(&invitation).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
		return
	}
	c.JSON(http.StatusOK, invitation)
}

func generateRegistrationCode() string {
	rand.Seed(time.Now().UnixNano())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func generateQRCode(data string) (string, error) {
	fileName := fmt.Sprintf("%s.png", data)
	err := qrcode.WriteFile(data, qrcode.Medium, 256, fileName)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func scheduleEmailVerification(email string) {
	time.Sleep(1 * time.Hour)
	sendEmailVerification(email)
}

func sendEmailVerification(to string) {
	from := os.Getenv("EMAIL_USERNAME")
	msg := fmt.Sprintf("Thank you for registering.")
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", fmt.Sprintf("Info Admin <%s>", from))
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", "Email Verification")
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
