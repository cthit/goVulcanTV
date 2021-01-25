package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/endpoints"
	"log"
)

var router *gin.Engine

func init() {
	log.Println("Initializing GIN webserver")
	router = gin.Default()

	api := router.Group("/api")
	{
		api.POST("/page_contents", endpoints.CreatePageContent)
		api.GET("/page_contents", endpoints.GetPageContents)
		api.DELETE("/page_contents/:id", endpoints.DeletePageContent)
		api.GET("/videos/next", endpoints.GetNext)
		api.GET("/videos/current", endpoints.GetCurrent)
		api.PUT("/videos/override/:id", endpoints.OverrideVideo)
	}
}

func Start() {
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver, err: %s\n", err)
	}
}