package routes

import (
	"github.com/ahmed-afzal1/go-auth/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("signup", controllers.SignUp)
}
