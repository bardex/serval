package web

import (
	"github.com/gin-gonic/gin"
)

func initRoutes(s *Server, router *gin.Engine) {
	router.GET("/", s.Home)

	// static files
	router.Static("/assets", "./web/assets/public")
	router.StaticFile("/favicon.ico", "./web/assets/public/favicon.ico")
}
