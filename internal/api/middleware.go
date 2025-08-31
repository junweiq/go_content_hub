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
	//TODO
	if sessionId == "" {
		ctx.AbortWithStatusJSON(http.StatusForbidden, "SID is null")
	}
	fmt.Println("SID = ", sessionId)
	ctx.Next()
}
