
package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"quill-backend/config"
	"quill-backend/database"
	"quill-backend/routes"
)

func main() {
	cfg := config.LoadConfig()

	db, err := database.ConnectDatabase(cfg.MongoURI)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	routes.SetupRoutes(router)

	log.Println("Quill server running on port", cfg.Port)

	router.Run(":" + cfg.Port)
}
