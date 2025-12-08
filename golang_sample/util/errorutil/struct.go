package errorutil

type ErrorMessage struct {
	Code            int         `json:"code,omitempty"`
	InternalCode    int         `json:"internal_code,omitempty"`
	Message         string      `json:"message"`
	InternalMessage string      `json:"internal_message"`
	ResponseData    interface{} `json:"response_data"`
}

type ErrorResponse struct {
	Message      string `json:"message"`
	InternalCode int    `json:"internal_code,omitempty"`
	TraceID      string `json:"trace_id"`
}

// IError - interface for error
type IError interface {
	error
	InternalError() string
	GetHTTPCode() int
	GetInternalCode() int
	GetResponseData() interface{}
	GetDataError() ErrorResponseDetail
}

type MetaDetail struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	InternalMsg string `json:"internal_msg,omitempty"`
}

type ErrorResponseDetail struct {
	Meta     MetaDetail  `json:"meta"`
	Response interface{} `json:"response"`
}

// Error - return message
func (b ErrorMessage) Error() string {
	return b.Message
}

// InternalError - return internal message
func (b ErrorMessage) InternalError() string {
	return b.InternalMessage
}

// GetHTTPCode - return code
func (b ErrorMessage) GetHTTPCode() int {
	return b.Code
}

// GetInternalCode - return code
func (b ErrorMessage) GetInternalCode() int {
	return b.InternalCode
}

// GetResponseData - return code
func (b ErrorMessage) GetResponseData() interface{} {
	return b.ResponseData
}

func (b ErrorMessage) GetDataError() ErrorResponseDetail {
	internalMsg := b.InternalError()
	response := b.GetResponseData()
	if response == nil {
		response = b.Error()
	}

	return ErrorResponseDetail{
		Meta: MetaDetail{
			Code:        b.GetInternalCode(),
			Msg:         b.Error(),
			InternalMsg: internalMsg,
		},
		Response: response,
	}
}
