package server

import (
	"context"
	"time"

	"gpaydemoopenapi/layer/interactor"
	"gpaydemoopenapi/pkg/config"
	"gpaydemoopenapi/util"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Service -
type Service struct {
	interactor *interactor.Interactor
	config     *config.Config
}

// NewHandler -
func NewHandler(interactor *interactor.Interactor) *Service {
	return &Service{interactor: interactor, config: interactor.App.Config}
}

// StartServer -
func (s *Service) StartServer() {
	e := echo.New()
	e.Use(s.interactor.LogCollect())
	if e == nil {
		panic("error start server")
	}
	routes := e.Group("/api/v1")
	routes.Use(middleware.Recover())
	// fund transfers
	routes.GET("/fund-to-bank", s.FundToBank)
	routes.GET("/information", s.GetInformation)
	routes.GET("/inquiry-account-bank", s.InquiryAccount)
	routes.GET("/list-bank", s.ListBank)
	routes.GET("/transfer-detail/:transfer_type/:transaction_id", s.GetTransferDetail)
	routes.GET("/list-bank", s.ListBank)
	routes.POST("/create-va", s.CreateVA)
	routes.POST("/update-va", s.VaUpdate)
	routes.POST("/close-va", s.VaClose)
	routes.POST("/detail-va", s.VaDetail)
	routes.POST("/reopen-va", s.VaReOpen)
	routes.GET("/init-bill-payment-gateway", s.InitBillPaymentGateway)
	routes.POST("/get-bill", s.GetBillInformation)
	routes.POST("/webhook-va", s.WebhookVA)
	routes.POST("/webhook-payment-gateway", s.WebhookPaymentGateway)

	// start server
	go func() {
		if err := e.Start(s.config.Port); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	util.HandleGracefulShutdown(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	})
}
