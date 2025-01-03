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
		client.POST("/create_client_to_message", baseAPIClient.CreateClientToMessage)
		client.POST("/create_collect_to_message", baseAPIClient.CreateCollectToMessage)
	}
	router.Run(":8080")
}
