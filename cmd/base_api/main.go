package main

import (
	"github.com/gin-gonic/gin"

	baseAPIClient "SpiderIM/internal/base_api/client"
)

func main() {
	router := gin.New()

	client := router.Group("/client")
	{
		client.POST("/create_client", baseAPIClient.CreateClient)
	}
	router.Run(":8080")
}
