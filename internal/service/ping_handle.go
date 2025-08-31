package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *CmsApp) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (c *CmsApp) PrivatePing(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "private pong",
	})
}
