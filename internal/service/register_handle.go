package service

import (
	"fmt"
	"go_content_hub/internal/dao"
	"go_content_hub/internal/modal"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRes struct {
}

func (c *CmsApp) Register(ctx *gin.Context) {
	var req RegisterReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDao := dao.NewUserDao(c.db)
	isExist, err := userDao.CheckExist(req.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if isExist {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "帳號已存在"})
		return
	}

	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nowTime := time.Now()
	if err := userDao.Create(&modal.User{
		Username:  req.Username,
		Password:  hashedPassword,
		Nickname:  req.Nickname,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password error = %v", err)
		return "", err
	}
	return string(hashedPassword), nil
}
