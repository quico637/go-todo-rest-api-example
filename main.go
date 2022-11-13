package main

import (
	"github.com/quico637/go-todo-rest-api-example/app"
	// "github.com/quico637/go-todo-rest-api-example/config"
	// "./app"
)

func main() {

	app := &app.App{}
	app.Initialize()
	app.Run(":3000")
}
