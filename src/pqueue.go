package main

import (
	"api"
)

// main starts http server
func main() {
	var app = api.App{}
	app.StartServer()
}
