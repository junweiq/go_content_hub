package service

import (
	"fmt"
	"go_content_hub/internal/dao"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRes struct {
	Sid      string `json:"sid"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
}

func (c *CmsApp) Login(ctx *gin.Context) {
	var req LoginReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userDao := dao.NewUserDao(c.db)
	user, err := userDao.FirstByUsername(req.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "帳號不存在"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "帳密不正確"})
		return
	}

	sid, err := c.generateSid(ctx, req.Username)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": &LoginRes{
			Sid:      sid,
			Username: user.Username,
			Nickname: user.Nickname,
		},
	})
}

func (c *CmsApp) generateSid(ctx *gin.Context, username string) (string, error) {
	sid := uuid.New().String()

	usidKey := "usid:" + username
	err := c.rdb.Set(ctx, usidKey, sid, 8*time.Hour).Err()
	if err != nil {
		fmt.Printf("generateSid() %s error [%v]", usidKey, err)
		return "", err
	}

	//sid create at
	usidCaKey := "usid_ca:" + sid
	err = c.rdb.Set(ctx, usidCaKey, time.Now().Unix(), 8*time.Hour).Err()
	if err != nil {
		fmt.Printf("generateSid() %s error [%v]", usidCaKey, err)
		return "", err
	}

	fmt.Println(usidKey + ", " + usidCaKey)

	return sid, nil
}
