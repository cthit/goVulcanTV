package endpoints

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	apiCommon "github.com/swexbe/govulcantv/internal/api/common"
	"github.com/swexbe/govulcantv/internal/common"
	"github.com/swexbe/govulcantv/internal/db/queries"
	"github.com/swexbe/govulcantv/internal/player"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func OverrideVideo(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.JSON(400, apiCommon.Response{
			Success: false,
			Error: fmt.Sprintf("Invalid pageContent id %s", idString),
		})
		return
	}

	pageContent, err := queries.GetPageContentById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(404, apiCommon.Response{
				Success: false,
				Error: "No such page content",
			})
			return
		}

		log.Printf("Error: failed to retrieve pageContent, err: %s\n", err)
		c.JSON(500, apiCommon.Response{
			Success: false,
			Error: "Unknown error",
		})
		return
	}

	player.ForcePlayVideo(&common.Video{
		Id:            pageContent.YoutubeID,
		LengthSeconds: pageContent.LengthSeconds,
	})

	c.JSON(200, apiCommon.Response{
		Success: true,
		Data: nil,
	})
}
