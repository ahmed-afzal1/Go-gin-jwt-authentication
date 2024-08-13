package routes

import (
	"github.com/ahmed-afzal1/go-auth/controllers"
	"github.com/ahmed-afzal1/go-auth/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("signup", controllers.SignUp)

	userRoute := r.Group("/users")

	userRoute.Use(middleware.Authenticate)
	{
		userRoute.GET("details", controllers.UserDetails)
	}
}
