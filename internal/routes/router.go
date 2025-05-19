package routes

import (
	"api-blog-go/internal/config"
	"api-blog-go/internal/handler"
	"api-blog-go/internal/middleware"
	"api-blog-go/internal/repository"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// Khởi tạo router
	r := gin.Default()

	// Khởi tạo repository và handler
	blogRepo := repository.NewBlogRepository(config.DB)
	blogHandler := handler.NewBlogHandler(blogRepo)

	userRepo := repository.NewUserRepository(config.DB)
	authHandler := handler.NewAuthHandler(userRepo)
	// Routes
	api := r.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}
		// Public routes
		blogs := api.Group("/blogs")
		{
			blogs.GET("", blogHandler.GetBlogs)
			blogs.GET("/:id", blogHandler.GetBlog)
		}

		// Protected routes
		protected := api.Group("/blogs")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.POST("", blogHandler.CreateBlog)
			protected.PUT("/:id", blogHandler.UpdateBlog)
			protected.DELETE("/:id", blogHandler.DeleteBlog)
		}

		// Ở đây bạn có thể thêm các routes khác trong tương lai
		// Ví dụ: users, categories, tags, ...
	}

	return r
}
