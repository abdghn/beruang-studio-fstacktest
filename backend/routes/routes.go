package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	adminCtrl := controllers.AdminController{DB: db}
	userCtrl := controllers.UserController{DB: db}

	r.POST("/admin/send-invitation", adminCtrl.SendInvitation)
	r.POST("/user/register", userCtrl.Register)
	r.GET("/user/invitation/:link", userCtrl.GetInvitation)

	return r
}
