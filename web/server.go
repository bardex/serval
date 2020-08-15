package web

import (
	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr/v2"
	"html/template"
)

type Server struct {
	router  *gin.Engine
	Version string
}

func (s *Server) Start() {
	s.router.Run()
}

func (s *Server) Create(releaseMode string) {
	gin.SetMode(releaseMode)
	s.router = gin.Default()
	s.loadTemplates()
	initRoutes(s, s.router)
}

func (s *Server) loadTemplates() {
	if gin.IsDebugging() {
		s.router.LoadHTMLGlob("web/templates/*")
	} else {
		tplBox := packr.New("templates", "templates")
		tpl := template.New("")
		tplBox.Walk(func(path string, f packd.File) error {
			tmpl := tpl.New(f.Name())
			data, _ := tplBox.FindString(f.Name())
			tmpl.Parse(data)
			return nil
		})
		s.router.SetHTMLTemplate(tpl)
	}
}
