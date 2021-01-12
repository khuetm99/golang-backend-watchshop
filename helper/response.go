package helper

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"timewise/model"
)

func ResponseErr(c echo.Context, code int, errMsg ...string) error {
	var msg string
	if len(errMsg) == 0 {
		msg = http.StatusText(code)
	} else {
		msg = errMsg[0]
	}
	return c.JSON(code, model.Response{
		StatusCode: code,
		Message:    msg,
	})
}

func ResponseData(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, model.Response{
		StatusCode: http.StatusOK,
		Message:    http.StatusText(http.StatusOK),
		Data:       data,
	})
}
