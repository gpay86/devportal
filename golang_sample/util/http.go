package util

import (
	"net/http"

	"gpaydemoopenapi/util/errorutil"

	"github.com/labstack/echo/v4"
)

type successResponse struct {
	Meta     Meta        `json:"meta"`
	Response interface{} `json:"response"`
}

type Meta struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	InternalMsg string `json:"internal_msg,omitempty"`
}

type errorResponse struct {
	Meta     Meta        `json:"meta"`
	Response interface{} `json:"response"`
}

// NeoErrorResponse - new group error response
func NeoErrorResponse(context echo.Context, err errorutil.IError) error {
	internalMsg := err.InternalError()

	response := err.GetResponseData()
	if response == nil {
		response = err.Error()
	}
	return context.JSON(200, errorResponse{
		Meta: Meta{
			Code:        err.GetInternalCode(),
			Msg:         err.Error(),
			InternalMsg: internalMsg,
		},
		Response: response,
	})
}

// NeoSuccessResponse - new group success response
func NeoSuccessResponse(context echo.Context, data interface{}) error {
	internalMsg := http.StatusText(http.StatusOK)
	if data != nil {
		return context.JSON(http.StatusOK, successResponse{
			Meta: Meta{
				Code:        http.StatusOK,
				Msg:         http.StatusText(http.StatusOK),
				InternalMsg: internalMsg,
			},
			Response: data,
		})
	}
	return context.JSON(http.StatusNoContent, nil)
}
