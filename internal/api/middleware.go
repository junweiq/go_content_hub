package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const SessionKey = "SID"

type SessionAuth struct{}

func NewSessionAuth() *SessionAuth {
	return &SessionAuth{}
}

func (s *SessionAuth) Auth(ctx *gin.Context) {
	sessionId := ctx.GetHeader(SessionKey)
	//TODO SID 較驗
	if sessionId == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, fmt.Sprintf("%s is null", SessionKey))
	}
	fmt.Println(fmt.Sprintf("%s =", SessionKey), sessionId)
	ctx.Next()
}
