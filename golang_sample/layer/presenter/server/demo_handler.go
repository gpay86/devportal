package server

import (
	"github.com/labstack/echo/v4"
	"gpaydemoopenapi/util"
)

func (s *Service) FundToBank(c echo.Context) error {
	result, ierr := s.interactor.FundToBank(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) GetInformation(c echo.Context) error {
	result, ierr := s.interactor.GetInformation(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) InquiryAccount(c echo.Context) error {
	result, ierr := s.interactor.Inquiry(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) ListBank(c echo.Context) error {
	result, ierr := s.interactor.GetListBank(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) GetTransferDetail(c echo.Context) error {
	result, ierr := s.interactor.GetTransferDetail(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) CreateVA(c echo.Context) error {
	result, ierr := s.interactor.CreateVA(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) VaUpdate(c echo.Context) error {
	result, ierr := s.interactor.UpdateVA(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) VaClose(c echo.Context) error {
	result, ierr := s.interactor.CloseVA(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) VaDetail(c echo.Context) error {
	result, ierr := s.interactor.DetailVA(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) VaReOpen(c echo.Context) error {
	result, ierr := s.interactor.ReOpenVA(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) InitBillPaymentGateway(c echo.Context) error {
	result, ierr := s.interactor.InitBillPaymentGateway(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) GetBillInformation(c echo.Context) error {
	result, ierr := s.interactor.GetBill(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) WebhookVA(c echo.Context) error {
	result, ierr := s.interactor.WebhookVA(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}

func (s *Service) WebhookPaymentGateway(c echo.Context) error {
	result, ierr := s.interactor.WebhookPaymentGateway(c)
	if ierr != nil {
		return util.NeoErrorResponse(c, ierr)
	}
	return util.NeoSuccessResponse(c, result)
}
