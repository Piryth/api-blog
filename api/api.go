package api

import (
	"github.com/gin-gonic/gin"
	"piryth.fr/blog/api/handlers"
	"piryth.fr/blog/database"
)

func SetupRoutes(r *gin.Engine, queries *database.Queries) {
	// Initialize the PostHandler
	postHandler := &handlers.PostHandler{Queries: queries}
	userHandler := &handlers.UserHandler{Queries: queries}

	// Define the routes
	r.POST("/posts", postHandler.CreatePost)
	r.GET("/posts/:id", postHandler.GetPost)
	r.GET("/posts", postHandler.ListPosts)
	r.PUT("/posts/:id", postHandler.UpdatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUser)
	r.GET("/users", userHandler.ListUsers)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)
}
