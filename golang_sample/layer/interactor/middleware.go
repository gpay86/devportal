package interactor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/micro/go-micro/v2/metadata"
	"gpaydemoopenapi/util"
	"io"
)

// LogCollect middleware
func (i Interactor) LogCollect() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bodyBytes, _ := io.ReadAll(c.Request().Body)
			traceID := util.BusinessID("DEMOAPI")
			c.Set("trace_id", traceID)
			checkMap := map[string]interface{}{}
			logData := bodyBytes
			json.Unmarshal(logData, &checkMap)
			i.Logger.InfoBody(traceID, c.Request().Method, fmt.Sprint(c.Request().URL), fmt.Sprint(c.RealIP()),
				string(logData), c.Request().Header)
			c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			ctx := metadata.Set(c.Request().Context(), "trace_id", traceID)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
