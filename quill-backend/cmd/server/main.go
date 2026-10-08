
package main

import (
	"log"

	"quill-backend/config"
	"quill-backend/database"
	"quill-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDatabase(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router, db)

	log.Println("Quill server running on port", cfg.Port)
	router.Run(":" + cfg.Port)
}

