package server

import (
	"github.com/gin-gonic/gin"
	"gobookcabin/gobookcabin"
	"net/http"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", gobookcabin.AppConfigurationInstance.ServerCorsAllowedOrigins)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Set-Cookie, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
