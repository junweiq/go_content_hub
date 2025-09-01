package api

import (
	"go_content_hub/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	authApiPrefix = "_api"
	apiPrefix     = "api"
)

func CmsRouter(r *gin.Engine) {
	cmsApp := service.CmsApp{}
	session := NewSessionAuth()
	api := r.Group(apiPrefix)
	{
		api.GET("ping", cmsApp.Ping)
		api.GET("cms/register", cmsApp.Register)
	}

	authApi := r.Group(authApiPrefix).Use(session.Auth)
	{
		authApi.GET("cms/hello", cmsApp.Hello)
	}
}
