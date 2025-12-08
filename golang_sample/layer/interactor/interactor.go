package interactor

import (
	"gpaydemoopenapi/layer/repository"
	gpayOpenapi "gpaydemoopenapi/layer/repository/gpay_openapi"
	redisRepo "gpaydemoopenapi/layer/repository/redis"
	"gpaydemoopenapi/pkg/logger"
	"gpaydemoopenapi/util/app"
)

// Interactor -
type Interactor struct {
	App                 *app.App
	Logger              *logger.Logger
	IRedis              repository.IRedis
	IServiceGPAYOpenAPI repository.IServiceGPAYOpenAPI
}

// Init -
func Init(app *app.App) *Interactor {

	i := &Interactor{
		App:                 app,
		Logger:              app.Logger,
		IRedis:              redisRepo.NewRedisRepositoryImpl(redisRepo.Connection(app.Config.Redis.Address, app.Config.Redis.Password, app.Config.Redis.DB)),
		IServiceGPAYOpenAPI: gpayOpenapi.NewServiceGpayOpenAPI(app.Config.Services.GPAYOpenAPI, app.Logger, app.Config),
	}
	return i
}
