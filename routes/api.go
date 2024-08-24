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

	postRoute := r.Group("/posts")

	postRoute.Use(middleware.Authenticate)
	{
		postRoute.GET("index", controllers.GetAllPost)
		postRoute.POST("create", controllers.CreatePost)
		postRoute.GET("/:id", controllers.FindPost)
	}

	categoryRoute := r.Group("/category")

	categoryRoute.Use(middleware.Authenticate)
	{
		categoryRoute.GET("/index", controllers.GetAllCategories)
		categoryRoute.POST("/create", controllers.CreateCategory)
	}
}
