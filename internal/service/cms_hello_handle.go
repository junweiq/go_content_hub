package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HelloReq struct {
	Name string `json:"name" binding:"required"`
}

type HelloRes struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) Hello(ctx *gin.Context) {
	var req HelloReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &HelloRes{
			Message: fmt.Sprintf("hello %s", req.Name),
		},
	})
}
