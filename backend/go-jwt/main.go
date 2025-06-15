package main

import (
	"go-jwt/controllers"
	"go-jwt/initializers"
	"go-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDB()
}
func main() {

	r := gin.Default()

	// Auth Routes
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	// Public Crop Routes
	r.GET("/crops", controllers.Cropindex)
	r.GET("/crops/:id", controllers.CropsShow)

	// Protected Crop Routes
	auth := r.Group("/")
	auth.Use(middleware.RequireAuth)
	{
		auth.POST("/crops", controllers.CropsCreate)
		auth.PUT("/crops/:id", controllers.CropsUpdate)
		auth.DELETE("/crops/:id", controllers.CropsDelete)
		auth.GET("/mycrops", controllers.MyCrops)
		auth.POST("/cart", controllers.AddToCart)
	}
	r.Run()
}
