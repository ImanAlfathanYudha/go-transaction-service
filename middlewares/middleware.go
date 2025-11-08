package middlewares

import (
	"go-transaction-service/common/response"
	errConstant "go-transaction-service/constants/error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func HandlePanic() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				logrus.Error("Recovered from panic:%v", r)
				c.JSON(http.StatusInternalServerError, response.Response{
					Status: errConstant.ErrInternalServerError.Error(),
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
