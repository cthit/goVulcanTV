package api

import (
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/endpoints"
	"log"
	"os"
)

var router *gin.Engine

func Init() {
	log.Println("Initializing GIN webserver")
	router = gin.Default()

	username := os.Getenv("admin_user")
	password := os.Getenv("admin_password")

	if username == "" {
		username = "digit"
	}
	if password == "" {
		password = "password"
	}

	log.Printf("Admin username: %s, pass: %s\n", username, password)

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

	err := router.Run()
	if err != nil {
		log.Fatalf("Failed to start webserver, err: %s\n", err)
	}
}