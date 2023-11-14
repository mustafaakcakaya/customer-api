package middlewares

import (
	"CustomerAPI/pkg/errors"
	//"github.com/sirupsen/logrus"
	"net/http"
)

func PanicExceptionHandling() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					//go logrus.Error(err)
					switch v := err.(type) {
					case *errors.Error:
						c.JSON(v.StatusCode, v.Public)
					case errors.Error:
						c.JSON(v.StatusCode, v.Public)
					default:
						c.NoContent(http.StatusInternalServerError)
					}
				}
			}()
			return next(c)
		}
	}
}
