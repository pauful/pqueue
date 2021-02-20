package main

import (
	"github.com/pauful/pqueue/src/api"
)

// main starts http server
func main() {
	var app = api.App{}
	app.StartServer()
}
