package response

import (
	"go-transaction-service/constants"
	errConstant "go-transaction-service/constants/error"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int
	Status  string
	Message any
	Data    interface{}
	Token   *string
	Notes   string
}

type ParamHTTPResp struct {
	Code    int
	Err     error
	Message *string
	Gin     *gin.Context
	Data    interface{}
	Token   *string
	Notes   *string
}

func HttpResponse(param ParamHTTPResp) {
	notes := ""
	if param.Err == nil {
		if param.Notes != nil {
			notes = *param.Notes
		}
		param.Gin.JSON(param.Code, Response{
			Code:    param.Code,
			Status:  constants.Success,
			Message: http.StatusText(http.StatusOK),
			Data:    param.Data,
			Token:   param.Token,
			Notes:   notes,
		})
		return
	}
	message := errConstant.ErrInternalServerError.Error()
	if param.Message != nil {
		message = *param.Message
	} else if param.Err != nil {
		if errConstant.ErrMapping(param.Err) {
			message = param.Err.Error()
		}
	}
	param.Gin.JSON(param.Code, Response{
		Code:    param.Code,
		Status:  constants.Error,
		Message: message,
		Data:    param.Data,
	})
	return
}
