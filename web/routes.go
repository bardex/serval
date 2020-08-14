package web

import (
	"github.com/gin-gonic/gin"
	"serval/web/actions"
)

func initRoutes(router *gin.Engine) {
	router.GET("/", actions.Home)

	// static files
	router.Static("/assets", "./web/assets/public")
	router.StaticFile("/favicon.ico", "./web/assets/public/favicon.ico")
}
