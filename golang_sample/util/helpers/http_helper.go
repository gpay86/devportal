package helpers

import (
	"github.com/labstack/echo/v4"
)

// BindingBody -
func BindingBody(body interface{}, context echo.Context) error {
	if err := context.Bind(body); err != nil {
		return err
	}

	if err := context.Validate(body); err != nil {
		return err
	}

	return nil
}
