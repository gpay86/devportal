package interactor

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gpaydemoopenapi/layer/model/aggregate/request/gpayopenapi_req"
	"gpaydemoopenapi/util/crypt"
	"gpaydemoopenapi/util/errorutil"
	"gpaydemoopenapi/util/helpers"
	"net/http"
	"strings"
	"time"
)

const (
	ACCESS_TOKEN_KEY string = "ACCESS_TOKEN_KEY_GPAY"
)

// FundToBank
func (i *Interactor) FundToBank(c echo.Context) (interface{}, errorutil.IError) {
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	transfer, err := i.IServiceGPAYOpenAPI.FundTransfer(c.Request().Context(), gpayopenapi_req.RequestTransfer2Bank{
		AccountNumber: "0687041020972",
		Amount:        10000,
		BankCode:      "VCCB",
		FullName:      "NGUYEN VAN A",
		Message:       "chuyen tien",
		TransactionId: uuid.New().String(),
		Type:          "ACCOUNT_NUMBER",
	}, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if transfer != nil {
		if transfer.Meta.Code.IsSuccess() {
			return transfer.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(transfer.Meta.Code),
			InternalCode:    cast.ToInt(transfer.Meta.Code),
			Message:         transfer.Meta.Msg,
			InternalMessage: transfer.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// GetInformation
func (i *Interactor) GetInformation(c echo.Context) (interface{}, errorutil.IError) {
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	information, err := i.IServiceGPAYOpenAPI.GetBalance(c.Request().Context(), token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if information != nil {
		if information.Meta.Code.IsSuccess() {
			return information.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(information.Meta.Code),
			InternalCode:    cast.ToInt(information.Meta.Code),
			Message:         information.Meta.Msg,
			InternalMessage: information.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// Inquiry
func (i *Interactor) Inquiry(c echo.Context) (interface{}, errorutil.IError) {
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	transfer, err := i.IServiceGPAYOpenAPI.InquiryAccountBank(c.Request().Context(), gpayopenapi_req.RequestInquiry{
		AccountNumber: "0687041020972",
		BankCode:      "VCCB",
		Type:          "ACCOUNT_NUMBER",
		RequestId:     uuid.New().String(),
	}, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if transfer != nil {
		if transfer.Meta.Code.IsSuccess() {
			return transfer.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(transfer.Meta.Code),
			InternalCode:    cast.ToInt(transfer.Meta.Code),
			Message:         transfer.Meta.Msg,
			InternalMessage: transfer.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// GetListBank
func (i *Interactor) GetListBank(c echo.Context) (interface{}, errorutil.IError) {
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.GetListBankIbft(c.Request().Context(), token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// GetTransferDetail
func (i *Interactor) GetTransferDetail(c echo.Context) (interface{}, errorutil.IError) {
	transferType := strings.TrimSpace(c.Param("transfer_type"))
	id := strings.TrimSpace(c.Param("transaction_id"))
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.TransferDetail(c.Request().Context(), transferType, id, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// getToken
func (i *Interactor) getToken(c echo.Context) (string, errorutil.IError) {
	token := i.IRedis.Get(ACCESS_TOKEN_KEY)
	if token == "" {
		tokenObj, err := i.IServiceGPAYOpenAPI.Login(c.Request().Context(), gpayopenapi_req.RequestLogin{
			ClientId:     i.App.Config.LoginInformation.ClientId,
			ClientSecret: i.App.Config.LoginInformation.ClientSecret,
		})
		if err != nil {
			return "", errorutil.ErrorMessage{
				Code:            http.StatusForbidden,
				InternalCode:    http.StatusForbidden,
				Message:         err.Error(),
				InternalMessage: err.Error(),
			}
		}
		expiredTime := tokenObj.Data.ExpiresIn - 10
		i.IRedis.Set(ACCESS_TOKEN_KEY, tokenObj.Data.AccessToken, time.Duration(expiredTime)*time.Second)
		token = tokenObj.Data.AccessToken
	}
	return token, nil
}

// CreateVA
func (i *Interactor) CreateVA(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.RequestCreateVA)
	helpers.BindingBody(body, c)
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.CreateVA(c.Request().Context(), *body, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// UpdateVA
func (i *Interactor) UpdateVA(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.UpdateVA)
	helpers.BindingBody(body, c)
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.UpdateVA(c.Request().Context(), *body, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// CloseVA
func (i *Interactor) CloseVA(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.CloseVA)
	helpers.BindingBody(body, c)
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.CloseVA(c.Request().Context(), *body, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// CloseVA
func (i *Interactor) DetailVA(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.CloseVA)
	helpers.BindingBody(body, c)
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.DetailVA(c.Request().Context(), *body, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// ReOpenVA
func (i *Interactor) ReOpenVA(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.CloseVA)
	helpers.BindingBody(body, c)
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.ReOpenVA(c.Request().Context(), *body, token)
	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// InitBillPaymentGateway
func (i *Interactor) InitBillPaymentGateway(c echo.Context) (interface{}, errorutil.IError) {
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.InitBill(c.Request().Context(), gpayopenapi_req.InitBill{
		Address:       "Testing",
		Amount:        15000,
		CallbackUrl:   "https://webhook.site/...",
		CustomerId:    "123456789",
		CustomerName:  "NGUYEN VAN A",
		Description:   "Thanh toan cho don hang 123",
		Email:         "nguyenvana@gmail.com",
		EmbedData:     "private data",
		PaymentMethod: "",
		PaymentType:   "IMMEDIATE",
		Phone:         "0123456789",
		RequestId:     uuid.New().String(),
		Title:         "Thanh toan cho don hang 123",
		WebhookUrl:    "https://webhook.site/...",
	}, token)

	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// GetBill
func (i *Interactor) GetBill(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.GetBill)
	helpers.BindingBody(body, c)
	token, errToken := i.getToken(c)
	if errToken != nil {
		return nil, errToken
	}
	banks, err := i.IServiceGPAYOpenAPI.GetBillInformation(c.Request().Context(), *body, token)

	if err != nil {
		return nil, errorutil.ErrorMessage{
			Code:            http.StatusBadRequest,
			InternalCode:    http.StatusBadRequest,
			Message:         err.Error(),
			InternalMessage: err.Error(),
		}
	}
	if banks != nil {
		if banks.Meta.Code.IsSuccess() {
			return banks.Data, nil
		}
		return nil, errorutil.ErrorMessage{
			Code:            cast.ToInt(banks.Meta.Code),
			InternalCode:    cast.ToInt(banks.Meta.Code),
			Message:         banks.Meta.Msg,
			InternalMessage: banks.Meta.InternalMsg,
		}
	}
	return nil, nil
}

// WebhookVA
func (i *Interactor) WebhookVA(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.WebhookVARequest)
	helpers.BindingBody(body, c)

	if body.Action == "CHANGE_BALANCE" {
		msgSignature := strings.Join([]string{
			fmt.Sprintf("gpay_trans_id=%v", body.GpayTransId),
			fmt.Sprintf("bank_trace_id=%v", body.BankTraceId),
			fmt.Sprintf("bank_transaction_id=%v", body.BankTransactionId),
			fmt.Sprintf("account_number=%v", body.AccountNumber),
			fmt.Sprintf("amount=%v", body.Amount),
			fmt.Sprintf("message=%v", body.Message),
			fmt.Sprintf("action=%v", "CHANGE_BALANCE"),
		}, "&")
		errVeriy := i.verifySignature(msgSignature, body.Signature)
		if errVeriy != nil {
			return nil, errVeriy
		}
		// business flow

	}
	if body.Action == "CLOSE_ACCOUNT" {
		msgSignature := strings.Join([]string{
			fmt.Sprintf("account_number=%v", body.AccountNumber),
			fmt.Sprintf("message=%v", body.Message),
			fmt.Sprintf("action=%v", "CLOSE_ACCOUNT"),
		}, "&")
		errVeriy := i.verifySignature(msgSignature, body.Signature)
		if errVeriy != nil {
			return nil, errVeriy
		}
		// business flow
	}
	// return http code 200 if handle success
	return nil, nil
}

// WebhookVA
func (i *Interactor) WebhookPaymentGateway(c echo.Context) (interface{}, errorutil.IError) {
	body := new(gpayopenapi_req.WebhookPaymentGatewayRequest)
	helpers.BindingBody(body, c)
	msgSignature := strings.Join([]string{
		fmt.Sprintf("merchant_order_id=%v", body.MerchantOrderID),
		fmt.Sprintf("gpay_trans_id=%v", body.GpayTransId),
		fmt.Sprintf("gpay_bill_id=%v", body.GpayBillId),
		fmt.Sprintf("status=%v", body.Status),
		fmt.Sprintf("embed_data=%v", body.EmbedData),
		fmt.Sprintf("user_payment_method=%v", body.UserPaymentMethod),
	}, "&")
	errVeriy := i.verifySignature(msgSignature, body.Signature)
	if errVeriy != nil {
		return nil, errVeriy
	}
	// business flow

	// return http code 200 if handle success
	return nil, nil
}

func (i *Interactor) verifySignature(msgSignature, signature string) errorutil.IError {
	err := crypt.VerifySHA256RSASign(msgSignature, signature)
	if err != nil {
		return errorutil.ErrorMessage{
			Code:         400,
			InternalCode: 400,
			Message:      "Unauthorized, signature doesn't matching",
			InternalMessage: fmt.Sprintf("Unauthorized, signature doesn't matching, msgSignature %v, signature: %v",
				msgSignature, signature),
		}
	}
	return nil
}
