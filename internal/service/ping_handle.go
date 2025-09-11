package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *CmsApp) PingHandle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
