package main

import (
	"api-blog-go/internal/config"
	"api-blog-go/internal/models"
	"api-blog-go/internal/routes"
)

func main() {
	// Khởi tạo database
	config.InitDatabase()

	// Auto Migrate
	config.DB.AutoMigrate(&models.Blog{}, &models.User{})

	// Setup router
	r := routes.SetupRouter()

	// Chạy server
	r.Run(":8080")
}
