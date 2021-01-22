package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"log"
)

func GetPageContents(c *gin.Context) {
	pageContents, err := queries.GetPageContents()
	if err != nil {
		log.Printf("Failed to retrieve page contents, err: %s\n", err)
		c.JSON(500, common.Response{
			Success: false,
			Error: "Failed to retrieve page contents",
		})
		return
	}
	c.JSON(200, common.Response{
		Success: true,
		Data: pageContents,
	})
}