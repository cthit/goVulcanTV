package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/endpoints"
	"log"
	"os"
)

var router *gin.Engine

func init() {
	log.Println("Initializing GIN webserver")
	router = gin.Default()

	username := os.Getenv("reset_db")
	password := os.Getenv("reset_db")

	if username == "" {
		username = "digit"
	}
	if password == "" {
		password = "password"
	}

	api := router.Group("/api")
	{

		auth := api.Group("/", gin.BasicAuth(gin.Accounts{
			username: password,
		}))
		{
			auth.POST("/page_contents", endpoints.CreatePageContent)
			auth.DELETE("/page_contents/:id", endpoints.DeletePageContent)
			auth.PUT("/videos/override/:id", endpoints.OverrideVideo)
		}
		api.GET("/page_contents", endpoints.GetPageContents)
		api.GET("/videos/next", endpoints.GetNext)
		api.GET("/videos/current", endpoints.GetCurrent)
	}
}

func Start() {
	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver, err: %s\n", err)
	}
}