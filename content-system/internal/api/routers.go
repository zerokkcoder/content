package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/services"
)

const (
	rootPath   = "/api"
	noAuthPath = "/out/api"
)

func CmsRouters(r *gin.Engine) {
	cmsApp := services.NewCmsApp()
	session := NewSessionAuth()
	root := r.Group(rootPath).Use(session.Auth)
	{
		// /api/cms/ping
		root.GET("/cms/ping", cmsApp.Ping)
		// /api/cms/content/create
		root.POST("/cms/content/create", cmsApp.ContentCreate)
		// /api/cms/content/update
		root.POST("/cms/content/update", cmsApp.ContentUpdate)
		// /api/cms/content/delete
		root.POST("/cms/content/delete", cmsApp.ContentDelete)
		// /api/cms/content/find
		root.GET("/cms/content/find", cmsApp.ContentFind)
	}

	noAuth := r.Group(noAuthPath)
	{
		// /out/api/cms/register
		noAuth.POST("/cms/register", cmsApp.Register)
		// /out/api/cms/login
		noAuth.POST("/cms/login", cmsApp.Login)
	}
}
