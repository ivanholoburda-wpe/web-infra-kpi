package main

import (
	"api-service/internal/config"
	"api-service/internal/database"
	"api-service/internal/handlers"
	"api-service/internal/repository"
	"api-service/internal/services"
	"api-service/pkg/mq"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	conf := config.InitConfig()
	db := database.InitDatabase(conf)

	publisher, _ := mq.NewRabbitMQPublisher(conf.MqUrl)
	siteRepository := repository.NewSiteRepositoryGORM(db)
	siteService := services.NewSiteService(siteRepository, publisher)
	siteHandler := handlers.NewSiteHandler(siteService)

	router := gin.Default()
	router.Use(corsMiddleware())

	api := router.Group("/api/v1")
	{
		sites := api.Group("/sites")
		{
			sites.POST("", siteHandler.Create)
			sites.GET("", siteHandler.GetAll)
			sites.GET("/:id", siteHandler.GetByID)
			sites.DELETE("/:id", siteHandler.Delete)
		}
	}

	log.Println("Server is running on port 8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}