package main

import (
	"github.com/Iheanacho-ai/go-authentication.git/controllers"
	"github.com/Iheanacho-ai/go-authentication.git/initializers"
	"github.com/Iheanacho-ai/go-authentication.git/middleware"
	"github.com/gin-gonic/gin"
)

func init(){
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.POST("/singup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:8080
}