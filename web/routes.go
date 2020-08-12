package web

import "serval/web/actions"

func initRoutes() {
	router.GET("/", actions.Home)

	// static files
	router.Static("/assets", "./web/assets/public")
	router.StaticFile("/favicon.ico", "./web/assets/public/favicon.ico")
}
