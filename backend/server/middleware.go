package server

import (
	"github.com/gin-gonic/gin"
	"gobookcabin/gobookcabin"
	"net/http"
)

// ErrorHandlerMiddleware maps the gobookcabin error codes to http status codes
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Reference/Status
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			code := gobookcabin.ErrorCode(err)
			message := gobookcabin.ErrorMessage(err)

			status := http.StatusInternalServerError
			switch code {
			case gobookcabin.ErrCodeInvalid:
				status = http.StatusBadRequest
			case gobookcabin.ErrCodeNotFound:
				status = http.StatusNotFound
			case gobookcabin.ErrCodeUnauthorized:
				status = http.StatusUnauthorized
			}
			c.JSON(status, gin.H{"error": message})
			c.Abort()
		}
	}
}
