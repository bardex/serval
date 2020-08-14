package main

import (
	"serval/web"
)

var ReleaseMode = "debug"
var Version = "devel"

func main() {
	var ws web.Server
	ws.Version = Version
	ws.Create(ReleaseMode)
	ws.Start()
}
