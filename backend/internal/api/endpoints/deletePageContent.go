package endpoints

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	apiCommon "github.com/swexbe/govulcantv/internal/api/common"
	"github.com/swexbe/govulcantv/internal/process"
	"github.com/swexbe/govulcantv/internal/vulcanTvErrors"
	"log"
	"strconv"
)

func DeletePageContent(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.JSON(400, apiCommon.Response{
			Success: false,
			Error: fmt.Sprintf("Invalid pageContent id %s", idString),
		})
		return
	}

	err = process.DeletePageContent(id)
	if err != nil {
		if errors.Is(err, vulcanTvErrors.NoSuchPageContent) {
			c.JSON(404, apiCommon.Response{
				Success: false,
				Error: fmt.Sprintf("No pagecontent with id %s", idString),
			})
			return
		}

		log.Printf("Failed to delete page content due to %s", err)
		c.JSON(500, apiCommon.Response{
			Success: false,
			Error:   "Failed to delete page content",
		})
		return
	}

	c.JSON(200, apiCommon.Response{
		Success: true,
		Data:    nil,
	})
}