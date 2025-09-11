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
	session := NewAuthMiddleware(cmsApp.Rdb)
	api := r.Group(apiPrefix)
	{
		api.GET("ping", cmsApp.PingHandle)
		api.POST("cms/register", cmsApp.RegisterHandle)
		api.POST("cms/login", cmsApp.LoginHandle)
	}

	authApi := r.Group(authApiPrefix).Use(session.AuthMiddleware)
	{
		authApi.GET("cms/hello", cmsApp.HelloHandle)
		authApi.GET("cms/content/create", cmsApp.ContentCreateHandle)
	}
}
