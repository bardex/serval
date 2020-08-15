package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packr/v2"
)

func initRoutes(s *Server, router *gin.Engine) {
	router.GET("/", s.Home)

	// static files
	static := packr.New("static", "public")
	router.StaticFS("/assets", static)
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.FileFromFS("/favicon.ico", static)
	})
}
