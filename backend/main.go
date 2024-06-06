package main

import (
	"backend/models"
	"backend/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	godotenv.Load()
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.Invitation{}, &models.Registration{})

	r := routes.SetupRouter(db)
	r.Run(":8080")
}
