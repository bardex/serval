package main

import (
	"serval/web"
)

var ReleaseMode = "debug"

func main() {
	var ws web.Server
	ws.Create(ReleaseMode)
	ws.Start()
}
