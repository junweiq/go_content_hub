package api

import (
	"go_content_hub/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	apiPrefix = "api"
)

func CmsRouter(r *gin.Engine) {
	cmsApp := service.CmsApp{}
	session := NewSessionAuth()
	api := r.Group(apiPrefix)
	{
		api.GET("ping", cmsApp.Ping)
	}
	authApi := api.Use(session.Auth)
	{
		authApi.GET("privatePing", cmsApp.PrivatePing)
	}
	//cmsApi := api.Group("cms")
	//{
	//	cmsApi.POST("login", cmsApp.Ping)
	//}
}
