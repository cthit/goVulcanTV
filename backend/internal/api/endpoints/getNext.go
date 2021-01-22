package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/swexbe/govulcantv/internal/api/common"
	"github.com/swexbe/govulcantv/internal/process"
)

func GetNext(c *gin.Context) {
	next := process.GetNext()
	c.JSON(200, common.Response{
		Success: true,
		Data:    next,
	})
}
