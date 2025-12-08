package gpayopenapi_res

type ErrorCode string

type Meta struct {
	Code        ErrorCode   `json:"code"`
	Error       ErrorDetail `json:"error"`
	InternalMsg string      `json:"internal_msg"`
	Msg         string      `json:"msg"`
}

type ErrorDetail struct {
	ErrorCode   string `json:"error_code"`
	Message     string `json:"message"`
	Path        string `json:"path"`
	SourceError string `json:"source_error"`
	Url         string `json:"url"`
}

type Response struct {
	Meta Meta `json:"meta"`
}

func (status ErrorCode) IsSuccess() bool {
	return status == "200"
}

func (status ErrorCode) IsVerifying() bool {
	return status == "202"
}

func (status ErrorCode) IsFailed() bool {
	return !status.IsSuccess() && !status.IsVerifying()
}

type ResponseLogin struct {
	Meta Meta     `json:"meta"`
	Data LoginRes `json:"data"`
}

type LoginRes struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type ResponseWal2Bank struct {
	Meta Meta         `json:"meta"`
	Data DataWal2bank `json:"data"`
}
type DataWal2bank struct {
	TransferCreatedTime       string `json:"transfer_created_time"`
	TransferId                string `json:"transfer_id"`
	TransferStatus            string `json:"transfer_status"`
	TransferStatusUpdatedTime string `json:"transfer_status_updated_time"`
	FeeAmount                 int    `json:"fee_amount"`
}

type ResponseMerchantInformation struct {
	Meta Meta                    `json:"meta"`
	Data DataMerchantInformation `json:"data"`
}

type DataMerchantInformation struct {
	AmountCash        int `json:"amount_cash"`
	AmountMinimum     int `json:"amount_minimum"`
	AmountRefund      int `json:"amount_refund"`
	AmountRevenue     int `json:"amount_revenue"`
	CashInInformation struct {
		AccountName   string `json:"account_name"`
		AccountNumber string `json:"account_number"`
		BankCode      string `json:"bank_code"`
		BankName      string `json:"bank_name"`
		QrCode        string `json:"qr_code"`
		QrCodeImage   string `json:"qr_code_image"`
	} `json:"cash_in_information"`
}

type ResponseInquiryAccount struct {
	Meta Meta        `json:"meta"`
	Data DataInquiry `json:"data"`
}

type DataInquiry struct {
	FullName string `json:"full_name"`
	OrderId  string `json:"order_id"`
	Status   string `json:"status"`
}

type ResponseListBank struct {
	Meta Meta `json:"meta"`
	Data []struct {
		BankBin  string `json:"bank_bin"`
		BankCode string `json:"bank_code"`
		BankName string `json:"bank_name"`
		Logo     string `json:"logo"`
	} `json:"data"`
}

type ResponseTransferDetail struct {
	Meta Meta         `json:"meta"`
	Data DataWal2bank `json:"data"`
}

type ResponseCreateVA struct {
	Meta Meta `json:"meta"`
	Data struct {
		AccountName   string `json:"account_name"`   // name of account
		AccountNumber string `json:"account_number"` // number of account
		AccountType   string `json:"account_type"`   // account type: enum: O, M(ONETIME, MANYTIME)
		Balance       int    `json:"balance"`
		EqualAmount   int    `json:"equal_amount"` // transfer equal amount
		ExpireAt      string `json:"expire_at"`    // has value if status is CLOSE
		MaxAmount     int    `json:"max_amount"`
		MinAmount     int    `json:"min_amount"`
		QrCode        string `json:"qr_code"`       // qr string: format vietQR
		QrCodeImage   string `json:"qr_code_image"` //image base64
		StartAt       string `json:"start_at"`      // created time
		Status        string `json:"status"`        // status is OPEN or CLOSE
	} `json:"data"`
}

type ResponseInitBill struct {
	Meta Meta `json:"meta"`
	Data struct {
		BillId      string `json:"bill_id"`  // bill id of GPAY
		BillUrl     string `json:"bill_url"` //redirect to url
		ExpiredTime string `json:"expired_time"`
		RequestId   string `json:"request_id"` // request id of merchant
	} `json:"data"`
}

type ResponseGetBill struct {
	Meta Meta `json:"meta"`
	Data struct {
		EmbedData         string `json:"embed_data"`
		GpayBillId        string `json:"gpay_bill_id"`
		GpayTransId       string `json:"gpay_trans_id"`
		MerchantOrderId   string `json:"merchant_order_id"`
		Status            string `json:"status"`
		UserPaymentMethod string `json:"user_payment_method"`
	} `json:"data"`
}
