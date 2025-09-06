package service

import (
	"fmt"
	"go_content_hub/internal/dao"
	"go_content_hub/internal/util"
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

	userDao := dao.NewUserDao(c.Db)
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
	expiredTime := 8 * time.Hour
	sid := uuid.New().String()

	sk := util.GetUserSidKey(username)
	err := c.Rdb.Set(ctx, sk, sid, expiredTime).Err()
	if err != nil {
		fmt.Printf("generateSid() %s error [%v]", sk, err)
		return "", err
	}

	sck := util.GetUserSidCreateAtKey(sid)
	err = c.Rdb.Set(ctx, sck, time.Now().Unix(), expiredTime).Err()
	if err != nil {
		fmt.Printf("generateSid() %s error [%v]", sck, err)
		return "", err
	}

	fmt.Println(sk + ", " + sck)

	return sid, nil
}
