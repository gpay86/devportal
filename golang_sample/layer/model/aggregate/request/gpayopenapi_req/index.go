package gpayopenapi_req

type RequestLogin struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type RequestTransfer2Bank struct {
	AccountNumber string `json:"account_number"`
	Amount        int    `json:"amount"`
	BankCode      string `json:"bank_code"`
	FullName      string `json:"full_name"`
	Message       string `json:"message"`
	TransactionId string `json:"transaction_id"`
	Type          string `json:"type"`
}

type RequestInquiry struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
	RequestId     string `json:"request_id"`
	Type          string `json:"type"`
}

type RequestCreateVA struct {
	AccountName     string `json:"account_name"`
	AccountType     string `json:"account_type"`
	BankCode        string `json:"bank_code"`
	CustomerAddress string `json:"customer_address"`
	Description     string `json:"description"`
	EqualAmount     int64  `json:"equal_amount"`
	MapId           string `json:"map_id"`
	MapType         string `json:"map_type"`
	MaxAmount       int64  `json:"max_amount"`
	MinAmount       int64  `json:"min_amount"`
}

type UpdateVA struct {
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_number"`
	AccountType   string `json:"account_type"`
	EqualAmount   int64  `json:"equal_amount"`
	MapId         string `json:"map_id"`
	MaxAmount     int64  `json:"max_amount"`
	MinAmount     int64  `json:"min_amount"`
}

type CloseVA struct {
	AccountNumber string `json:"account_number"`
	CloseReason   string `json:"close_reason"`
}

type InitBill struct {
	Address       string `json:"address"`
	Amount        int64  `json:"amount"`
	CallbackUrl   string `json:"callback_url"`  // is redirect url when user finish payment, method GET
	CustomerId    string `json:"customer_id"`   // customer id of merchant
	CustomerName  string `json:"customer_name"` // customer name of merchant
	Description   string `json:"description"`   // description of transaction
	Email         string `json:"email"`
	EmbedData     string `json:"embed_data"`     // data of merhcant, return in callback and webhook event
	PaymentMethod string `json:"payment_method"` // method payment: BANK_ATM,BANK_INTERNATIONAL,VA. if use all method set empty
	PaymentType   string `json:"payment_type"`   // IMMEDIATE,DELAYED: fix is IMMEDIATE
	Phone         string `json:"phone"`          // phone no of customer
	RequestId     string `json:"request_id"`     // request id of merchant: unique
	Title         string `json:"title"`          // title of transaction
	WebhookUrl    string `json:"webhook_url"`    // method POST, send when transaction finish: ORDER_SUCCESS,ORDER_FAILED, ...
}

type GetBill struct {
	GpayBillId      string `json:"gpay_bill_id"`
	MerchantOrderId string `json:"merchant_order_id"`
}

type WebhookVARequest struct {
	GpayTransId       string `json:"gpay_trans_id"`
	BankTraceId       string `json:"bank_trace_id"`
	BankTransactionId string `json:"bank_transaction_id"`
	AccountNumber     string `json:"account_number"`
	Amount            uint64 `json:"amount"`
	Message           string `json:"message"`
	Action            string `json:"action"`
	MerchantCode      string `json:"merchant_code"`
	Signature         string `json:"signature"`
}

type WebhookPaymentGatewayRequest struct {
	MerchantOrderID   string `json:"merchant_order_id"`
	GpayTransId       string `json:"gpay_trans_id"`
	GpayBillId        string `json:"gpay_bill_id"`
	Status            string `json:"status"` // ORDER_SUCCESS, ORDER_FAILED
	EmbedData         string `json:"embed_data"`
	UserPaymentMethod string `json:"user_payment_method"` //BANK_ATM,BANK_INTERNATIONAL,VA
	Signature         string `json:"signature"`
}
