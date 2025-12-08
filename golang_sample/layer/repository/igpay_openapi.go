package repository

import (
	"context"
	"gpaydemoopenapi/layer/model/aggregate/request/gpayopenapi_req"
	"gpaydemoopenapi/layer/model/aggregate/response/gpayopenapi_res"
)

type IServiceGPAYOpenAPI interface {
	Login(ctx context.Context, req gpayopenapi_req.RequestLogin) (res *gpayopenapi_res.ResponseLogin, err error)
	FundTransfer(ctx context.Context, req gpayopenapi_req.RequestTransfer2Bank, token string) (res *gpayopenapi_res.ResponseWal2Bank, err error)
	GetBalance(ctx context.Context, token string) (res *gpayopenapi_res.ResponseMerchantInformation, err error)
	InquiryAccountBank(ctx context.Context, req gpayopenapi_req.RequestInquiry, token string) (res *gpayopenapi_res.ResponseInquiryAccount, err error)
	GetListBankIbft(ctx context.Context, token string) (res *gpayopenapi_res.ResponseListBank, err error)
	TransferDetail(ctx context.Context, typeTransfer, transactionID, token string) (res *gpayopenapi_res.ResponseTransferDetail, err error)
	CreateVA(ctx context.Context, req gpayopenapi_req.RequestCreateVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error)
	UpdateVA(ctx context.Context, req gpayopenapi_req.UpdateVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error)
	CloseVA(ctx context.Context, req gpayopenapi_req.CloseVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error)
	DetailVA(ctx context.Context, req gpayopenapi_req.CloseVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error)
	ReOpenVA(ctx context.Context, req gpayopenapi_req.CloseVA, token string) (res *gpayopenapi_res.ResponseCreateVA, err error)
	InitBill(ctx context.Context, req gpayopenapi_req.InitBill, token string) (res *gpayopenapi_res.ResponseInitBill, err error)
	GetBillInformation(ctx context.Context, req gpayopenapi_req.GetBill, token string) (res *gpayopenapi_res.ResponseGetBill, err error)
}
