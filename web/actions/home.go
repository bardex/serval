package actions

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {

	c.HTML(
		http.StatusOK,
		"home.html",
		gin.H{"Name": "Vasya"},
	)
}
