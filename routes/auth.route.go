package routes

import (
	"Wellness-monitoring/controllers"
	"Wellness-monitoring/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthRoutes struct {
	authController controllers.AuthController
	db             *gorm.DB
}

func NewAuthRoutes(authController controllers.AuthController, db *gorm.DB) AuthRoutes {
	return AuthRoutes{authController, db}
}

func (rc *AuthRoutes) AuthRoute(rg *gin.RouterGroup) {

	router := rg.Group("/auth")
	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(rc.db), rc.authController.LogoutUser)
}
