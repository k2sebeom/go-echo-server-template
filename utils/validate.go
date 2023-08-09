package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RequestValidator struct {
	validator *validator.Validate
}

func (v *RequestValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func NewRequestValidator() *RequestValidator {
	v := &RequestValidator{
		validator: validator.New(),
	}
	return v
}

func Validate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	return c.Validate(i)
}
