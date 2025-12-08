package main

import (
	"gpaydemoopenapi/layer/interactor"
	"gpaydemoopenapi/layer/presenter/server"
	"gpaydemoopenapi/util/app"
)

func main() {
	// init interactor
	app := app.InitApp(app.Params{
		ConfigPath: "./",
		ConfigName: "config",
	})
	interactor := interactor.Init(app)

	server.NewHandler(interactor).StartServer()
}
