package controller

import (
	"fmt"
	errWrap "go-transaction-service/common/error"
	"go-transaction-service/common/response"
	errConstant "go-transaction-service/constants/error"
	"go-transaction-service/domain/dto"
	"go-transaction-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service services.IServiceRegistry
}

type ITransactionController interface {
	UploadTransactionCSV(ctx *gin.Context)
	GetAllBalance(ctx *gin.Context)
	GetAllIssues(ctx *gin.Context)
}

func NewTransactionController(service services.IServiceRegistry) ITransactionController {
	return &TransactionController{service: service}
}

func (t *TransactionController) GetAllIssues(ctx *gin.Context) {
	issues, err := t.service.GetTransaction().GetAllIssues(ctx)
	errorMessage := ""
	if err != nil {
		errorMessage = fmt.Sprintf("%v", err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Message: &errorMessage,
			Err:     err,
			Gin:     ctx,
		})
		return
	}
	result := gin.H{
		"issues": issues,
	}
	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (t *TransactionController) GetAllBalance(ctx *gin.Context) {
	transactions, totalBalance, err := t.service.GetTransaction().GetAllBalance(ctx)
	errorMessage := ""
	if err != nil {
		errorMessage = fmt.Sprintf("%v", err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Message: &errorMessage,
			Err:     err,
			Gin:     ctx,
		})
		return
	}
	result := gin.H{
		"transactions":  transactions,
		"total_balance": totalBalance,
	}

	response.HttpResponse(response.ParamHTTPResp{
		Code: http.StatusOK,
		Data: result,
		Gin:  ctx,
	})
}

func (t *TransactionController) UploadTransactionCSV(ctx *gin.Context) {
	// --- Step 1: Retrieve file from request
	file, header, err := ctx.Request.FormFile("file")
	errorMessage := ""
	if err != nil {
		errorMessage := fmt.Sprintf("%v: %v", errWrap.WrapError(errConstant.ErrFileNotFound), "Content is empty")
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Message: &errorMessage,
			Err:     err,
			Gin:     ctx,
		})
		return
	}
	defer file.Close()

	// --- Step 2: Construct DTO
	req := dto.UploadRequest{
		File:   file,
		Header: header,
	}

	// --- Step 3: Validate file type
	if req.Header == nil || req.Header.Filename == "" {
		errorMessage = "File either empty or invalid"
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Message: &errorMessage,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	if ext := getFileExtension(req.Header.Filename); ext != "csv" {
		err = errConstant.ErrInvalidCSVFormat
		errorMessage = fmt.Sprintf("%v", err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Message: &errorMessage,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	// --- Step 4: Call service to process CSV
	err, errorDetails := t.service.GetTransaction().UploadTransactionCSV(ctx, req.File)
	if err != nil {
		errorMessage = fmt.Sprintf("%v", err)
		response.HttpResponse(response.ParamHTTPResp{
			Code:    http.StatusBadRequest,
			Message: &errorMessage,
			Err:     err,
			Gin:     ctx,
		})
		return
	}

	// --- Step 5: Return success response
	successMessage := "CSV uploaded successfully. "
	if len(errorDetails) > 0 {
		successMessage = successMessage + "Some rows were skipped due to format errors"
	}
	errorNotes := ""
	if len(errorDetails) > 0 {
		for i, e := range errorDetails {
			if i == 0 {
				errorNotes += e
			} else {
				errorNotes += ". " + e
			}
		}
	}
	response.HttpResponse(response.ParamHTTPResp{
		Code:    http.StatusOK,
		Message: &successMessage,
		Gin:     ctx,
		Notes:   &errorNotes,
	})
}

func getFileExtension(filename string) string {
	if len(filename) < 4 {
		return ""
	}
	return filename[len(filename)-3:]
}
