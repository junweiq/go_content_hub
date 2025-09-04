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
	cmsApp := service.NewCmsApp()
	session := NewSessionAuth()
	api := r.Group(apiPrefix)
	{
		api.GET("ping", cmsApp.Ping)
		api.POST("cms/register", cmsApp.Register)
		api.POST("cms/login", cmsApp.Login)
	}

	authApi := r.Group(authApiPrefix).Use(session.Auth)
	{
		authApi.GET("cms/hello", cmsApp.Hello)
	}
}
