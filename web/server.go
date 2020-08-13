package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"html/template"
)

var router *gin.Engine

func Start() {
	//gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	tplBox := packr.New("templates", "assets/templates")

	tpl := template.New("")

	tplBox.Walk(func(path string, f packd.File) error {
		tmpl := tpl.New(f.Name())
		data, _ := tplBox.FindString(f.Name())
		tmpl.Parse(data)

		return nil
	})

	router.SetHTMLTemplate(tpl)

	initRoutes()
	router.Run()
}
