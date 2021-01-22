package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/common"
	"github.com/swexbe/govulcantv/internal/process"
)

func GetCurrent(c *gin.Context) {
	c.JSON(200, common.Response{
		Success: true,
		Data:    process.GetCurrent(),
	})
}
