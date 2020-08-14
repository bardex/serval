package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) Home(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{"version": s.Version},
	)
}
