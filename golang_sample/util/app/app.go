package app

import (
	"os"

	"gpaydemoopenapi/pkg/config"
	"gpaydemoopenapi/pkg/logger"
)

// InitApp data
func InitApp(params Params) *App {
	// load config
	config, err := config.LoadConfig(params.ConfigPath, params.ConfigName)
	if err != nil {
		panic(err)
	}

	// init logger
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unkown"
	}
	env := config.ENV
	log, err := logger.NewLogger(hostname, env)
	if err != nil {
		panic(err.Error())
	}

	// init mongo

	return &App{
		Config: config,
		Logger: log,
	}
}
