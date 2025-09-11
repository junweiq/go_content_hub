package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ContentCreateHandleReq struct {
	Name string `json:"name" binding:"required"`
}

type ContentCreateHandleRes struct {
	Message string `json:"message" binding:"required"`
}

func (c *CmsApp) ContentCreateHandle(ctx *gin.Context) {
	var req ContentCreateHandleReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &ContentCreateHandleRes{
			Message: fmt.Sprintf("hello %s", req.Name),
		},
	})
}
