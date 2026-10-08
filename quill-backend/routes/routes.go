
package routes

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"

	"quill-backend/handlers"
)

func SetupRoutes(router *gin.Engine, db *mongo.Database) {
	router.POST("/api/notes", handlers.CreateNote)
	router.GET("/api/notes", handlers.GetNotes)
	router.GET("/api/notes/:id", handlers.GetNote)
	router.PUT("/api/notes/:id", handlers.UpdateNote)
	router.DELETE("/api/notes/:id", handlers.DeleteNote)

	router.POST("/api/auth/register", handlers.RegisterUser)
	router.POST("/api/auth/login", handlers.LoginUser)
	router.GET("/api/auth/me", handlers.GetUser)
	router.PUT("/api/auth/me", handlers.UpdateUser)
	router.DELETE("/api/auth/me", handlers.DeleteUser)

	router.POST("/api/tasks", handlers.CreateTask)
	router.GET("/api/tasks", handlers.GetTasks)
	router.GET("/api/tasks/:id", handlers.GetTask)
	router.PUT("/api/tasks/:id", handlers.UpdateTask)
	router.DELETE("/api/tasks/:id", handlers.DeleteTask)
}
