package main

import (
	"github.com/ahmed-afzal1/go-auth/config"
	_ "github.com/ahmed-afzal1/go-auth/docs"
	"github.com/ahmed-afzal1/go-auth/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Use(gin.Logger())

	config.ConnectDatabase()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r)

	r.Run()
}
