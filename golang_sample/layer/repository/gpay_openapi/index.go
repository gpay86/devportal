package gpay_openapi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"gpaydemoopenapi/layer/model/aggregate/request/gpayopenapi_req"
	"gpaydemoopenapi/layer/model/aggregate/response/gpayopenapi_res"
	"gpaydemoopenapi/layer/repository"
	"gpaydemoopenapi/pkg/config"
	"gpaydemoopenapi/pkg/logger"
	"gpaydemoopenapi/util"
	"gpaydemoopenapi/util/crypt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const timeout = time.Minute*3 + time.Second*5

type repoImpl struct {
	Uri    string
	Logger *logger.Logger
	Conf   *config.Config
}

type Response struct {
	Meta Res `json:"meta"`
}

type Res struct {
	Code        int    `json:"code"`
	InternalMsg string `json:"internal_msg"`
	Msg         string `json:"msg"`
}

func (this *repoImpl) Login(ctx context.Context, req gpayopenapi_req.RequestLogin) (res *gpayopenapi_res.ResponseLogin, err error) {
	headers := make(map[string]string)
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/auth/token", req, &res, headers)
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) FundTransfer(ctx context.Context, req gpayopenapi_req.RequestTransfer2Bank, token string) (res *gpayopenapi_res.ResponseWal2Bank, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/payouts/instant/transfer-to-bank", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) InquiryAccountBank(ctx context.Context, req gpayopenapi_req.RequestInquiry, token string) (res *gpayopenapi_res.ResponseInquiryAccount, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/payouts/bank-account/query", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) GetListBankIbft(ctx context.Context, token string) (res *gpayopenapi_res.ResponseListBank, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodGet, "/v1/reference/banks", nil, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) TransferDetail(ctx context.Context, typeTransfer, transactionID, token string) (res *gpayopenapi_res.ResponseTransferDetail, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodGet, fmt.Sprintf("/v1/reporting/transactions/%v/%v", typeTransfer, transactionID), nil, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) GetBalance(ctx context.Context, token string) (res *gpayopenapi_res.ResponseMerchantInformation, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodGet, "/v1/account/information", nil, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) CreateVA(ctx context.Context, req gpayopenapi_req.RequestCreateVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/collection/va/create", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) UpdateVA(ctx context.Context, req gpayopenapi_req.UpdateVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/collection/va/update", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) CloseVA(ctx context.Context, req gpayopenapi_req.CloseVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/collection/va/close", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) DetailVA(ctx context.Context, req gpayopenapi_req.CloseVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/collection/va/detail", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) ReOpenVA(ctx context.Context, req gpayopenapi_req.CloseVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/collection/va/re-open", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) InitBill(ctx context.Context, req gpayopenapi_req.InitBill, token string) (res *gpayopenapi_res.ResponseInitBill, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/payments/gateway/init-order", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) GetBillInformation(ctx context.Context, req gpayopenapi_req.GetBill, token string) (res *gpayopenapi_res.ResponseGetBill, err error) {
	err = this.sendRequestWithHeader(ctx, http.MethodPost, "/v1/payments/gateway/query-order", req, &res, this.GetHeader(token))
	if err != nil {
		return nil, err
	}
	return
}

func (this *repoImpl) GetHeader(token string) map[string]string {
	headers := make(map[string]string)
	headers["Authorization"] = fmt.Sprintf("Bearer %v", token)
	headers["x-certificate"] = crypt.GetCert()
	headers["x-requests-id"] = uuid.New().String()
	headers["x-timestamp"] = cast.ToString(time.Now().UnixNano() / 1e6)
	return headers
}

func (this *repoImpl) BuildSignature(headers map[string]string, request string) string {
	messgage := fmt.Sprintf("%v%v%v", headers["x-timestamp"], headers["x-requests-id"], request)
	signature := crypt.EncryptRSA(messgage)
	return signature
}

func (this *repoImpl) sendRequestWithHeader(ctx context.Context, method, path string, data interface{}, res interface{}, headers map[string]string) (err error) {
	resB, status, err := this.SendRequest(ctx, Request{
		Method:  method,
		Path:    path,
		Data:    data,
		Headers: headers,
	})
	if err != nil {
		return
	}
	var resp = new(gpayopenapi_res.Response)
	if status != http.StatusOK && status != http.StatusCreated {
		_ = json.Unmarshal(resB, resp)
		res = resp
		if resp.Meta.Msg != "" || resp.Meta.Msg != "" {
			messageInternal := resp.Meta.Error.Message
			if messageInternal == "" {
				messageInternal = resp.Meta.Msg
			}
			if messageInternal == "" {
				messageInternal = resp.Meta.Msg
			}
			messagePublic := resp.Meta.Msg
			if messagePublic == "" {
				messagePublic = resp.Meta.Msg
			}
			errStr := resp.Meta.Msg
			if errStr == "" {
				errStr = resp.Meta.Msg
			}
			internalCodeStr := resp.Meta.Error.ErrorCode
			if internalCodeStr == "" {
				internalCodeStr = fmt.Sprint(resp.Meta.Code)
			}

			internalCode, _ := strconv.Atoi(internalCodeStr)
			err = ErrorMsg{
				Code:         string(resp.Meta.Code),
				Message:      messagePublic,
				error:        fmt.Errorf(errStr),
				InternalMsg:  messageInternal,
				InternalCode: internalCode,
			}
		} else {
			err = ErrorMsg{
				Code:         string(status),
				Message:      "Có lỗi xảy ra.",
				error:        fmt.Errorf("Có lỗi xảy ra."),
				InternalMsg:  "Có lỗi xảy ra.",
				InternalCode: status,
			}
		}
		return
	}

	json.Unmarshal(resB, res)
	_ = json.Unmarshal(resB, resp)
	if string(resp.Meta.Code) != cast.ToString(http.StatusOK) && string(resp.Meta.Code) != cast.ToString(http.StatusCreated) {
		internalCode := status
		if string(resp.Meta.Code) == "0" || string(resp.Meta.Code) == "" {
			status = http.StatusInternalServerError
			internalCode = 99999
		}
		err = ErrorMsg{
			Code:         string(status),
			Message:      "Có lỗi xảy ra.",
			error:        fmt.Errorf("Có lỗi xảy ra."),
			InternalMsg:  "Có lỗi xảy ra.",
			InternalCode: internalCode,
		}
		return
	}
	return
}

// SendRequest -
func (this *repoImpl) SendRequest(ctx context.Context, r Request) (resB []byte, status int, err error) {
	// new request
	client := new(http.Client)
	client.Timeout = timeout
	b, err := json.Marshal(r.Data)
	if err != nil {
		return
	}
	req, err := http.NewRequest(r.Method, fmt.Sprintf("%v%v", this.Uri, r.Path), strings.NewReader(string(b)))
	if err != nil {
		this.Logger.Infof("send request error, url: %v, request: %v, err: %v", fmt.Sprintf("%v%v", this.Uri, r.Path), r, err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	if len(r.Headers) > 0 {
		for key, value := range r.Headers {
			req.Header.Add(key, value)
		}
		requestSign := string(b)
		if r.Method == http.MethodGet {
			requestSign = ""
		}
		req.Header.Add("signature", this.BuildSignature(r.Headers, requestSign))
	}

	req = req.WithContext(util.ContextWithTimeOut())
	req.Close = true

	// call
	res, err := client.Do(req)
	if err != nil {
		return
	}

	// get result bytes, status
	resB, err = ioutil.ReadAll(res.Body)
	this.Logger.Infof("path: %v res: %+v request %+v statusCode %+v, headers %v", fmt.Sprintf("%v%v", this.Uri, r.Path), string(resB), string(b), res.StatusCode, req.Header)
	defer res.Body.Close()
	status = res.StatusCode

	return
}

type ErrorMsg struct {
	Message      string
	Code         string
	InternalMsg  string
	InternalCode int
	OrderID      string
	error
}

type Request struct {
	Method                  string
	Path                    string
	Data                    interface{}
	Headers                 map[string]string
	DisabledVerifySignature bool
}

func NewServiceGpayOpenAPI(uri string, logger *logger.Logger, conf *config.Config) repository.IServiceGPAYOpenAPI {
	return &repoImpl{
		Uri:    uri,
		Logger: logger,
		Conf:   conf,
	}
}
