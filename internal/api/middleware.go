package api

import (
	"errors"
	"fmt"
	"go_content_hub/internal/constant"
	"go_content_hub/internal/util"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type SessionAuth struct {
	Rdb redis.Client
}

func NewSessionAuth(Rdb *redis.Client) *SessionAuth {
	return &SessionAuth{
		Rdb: *Rdb,
	}
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sid := ctx.GetHeader(constant.SessionKey)

	if sid == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, fmt.Sprintf("%s is null", constant.SessionKey))
		return
	}

	sck := util.GetUserSidCreateAtKey(sid)
	loginTime, err := s.Rdb.Get(ctx, sck).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, fmt.Sprintf("%s auth error", constant.SessionKey))
		return
	}
	if loginTime == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, fmt.Sprintf("%s auth fail", constant.SessionKey))
		return
	}

	ctx.Next()
}
