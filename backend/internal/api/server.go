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
		api.GET("/pageContents", endpoints.GetPageContents)
	}
}

func Start() {
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver, err: %s\n", err)
	}
}