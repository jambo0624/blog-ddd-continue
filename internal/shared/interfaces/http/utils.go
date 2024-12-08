package http

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseUintParam(c *gin.Context, param string) uint {
	id, err := strconv.ParseUint(c.Param(param), 10, 32)
	if err != nil {
		return 0
	}
	return uint(id)
}
