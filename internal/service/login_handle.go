package service

import (
	"go_content_hub/internal/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	SID      string `json:"sid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDao := dao.NewUserDao(c.db)
	user, err := userDao.FirstByUsername(req.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "帳號不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "帳密不正確"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LoginRes{
			SID:      generateSid(),
			Username: user.Username,
			Nickname: user.Nickname,
		},
	})
}

func generateSid() string {
	//TODO sid 生成
	//TODO sid 持久化
	return "123456"
}
