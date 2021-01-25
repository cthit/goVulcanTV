package endpoints

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/common"
	"github.com/swexbe/govulcantv/internal/db/models"
	"github.com/swexbe/govulcantv/internal/process"
	"github.com/swexbe/govulcantv/internal/vulcanTvErrors"
	"io/ioutil"
	"log"
)

type CreatePageContentResponse struct {
	ID uint64 `json:"ID"`
}

func CreatePageContent(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("Error: failed to read json data, err: %s\n", err)
		c.JSON(500, common.Response{
			Success: false,
			Error: "Failed to read request data",
		})
		return
	}

	var receivedPageContent models.PageContent
	err = json.Unmarshal(jsonData, &receivedPageContent)
	if err != nil {
		log.Printf("Error: invalid json data, err: %s\n", err)
		c.JSON(400, common.Response{
			Success: false,
			Error: "Invalid json data",
		})
		return
	}

	id, err := process.CreatePageContent(&receivedPageContent)
	if err != nil {
		log.Printf("Error: failed to create pageContent, err: %s\n", err)
		if errors.Is(err, vulcanTvErrors.ErrAlreadyExists) {
			c.JSON(200, common.Response{
				Success: false,
				Error: "This video & length already exists",
			})
			return
		}

		c.JSON(500, common.Response{
			Success: false,
			Error: "Failed to create pageContent",
		})
		return
	}

	c.JSON(201, common.Response{
		Success: true,
		Data: CreatePageContentResponse{
			ID: id,
		},
	})
}
