package app

import (
	"gpaydemoopenapi/pkg/config"
	"gpaydemoopenapi/pkg/logger"
)

// App data
type App struct {
	Config *config.Config
	Logger *logger.Logger
}

// Params app
type Params struct {
	ConfigPath string
	ConfigName string
}
