package main

import (
	"serval/web"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	// start web server
	wg.Add(1)
	go web.Start()

	// wait all goroutines
	wg.Wait()
}
