package main

import (
	"fmt"
	"log"

	"github.com/GuilhermePT1/api-social-meli/internal/domain/models"
	"github.com/GuilhermePT1/api-social-meli/internal/infra/database"
	"github.com/GuilhermePT1/api-social-meli/internal/infra/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database: " + err.Error())
	}

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Post{}, &models.Follow{})

	router := gin.Default()
	routes.SetupRoutes(router, db)

	router.Run(":8080")

	fmt.Println("API is running on port 8080")
	fmt.Println("Database connected successfully")
}
