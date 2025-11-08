package services

import (
	"context"
	"encoding/csv"
	"fmt"
	errWrap "go-transaction-service/common/error"
	errConstant "go-transaction-service/constants/error"
	"go-transaction-service/domain/model"
	repositories "go-transaction-service/repositories"
	"io"
	"strconv"
	"strings"
)

type ITransactionService interface {
	UploadTransactionCSV(context.Context, io.Reader) (error, []string)
	GetAllBalance(ctx context.Context) ([]model.Transaction, float64, error)
	GetAllIssues(ctx context.Context) ([]model.Transaction, error)
}

type TransactionService struct {
	respository repositories.IRepositoryRegistry
}

func NewTransactionService(respository repositories.IRepositoryRegistry) ITransactionService {
	return &TransactionService{respository: respository}
}

func (t *TransactionService) GetAllBalance(ctx context.Context) ([]model.Transaction, float64, error) {
	transactions, err := t.respository.GetTransaction().GetAllTransactions(ctx)
	if err != nil {
		return nil, 0, err
	}
	var totalCredits, totalDebits float64
	var successfulTxns []model.Transaction

	for _, txn := range transactions {
		if strings.ToLower(txn.Status) == "success" {
			successfulTxns = append(successfulTxns, txn)
			if strings.ToLower(txn.Type) == "credit" {
				totalCredits += txn.Amount
			} else if strings.ToLower(txn.Type) == "debit" {
				totalDebits += txn.Amount
			}
		}
	}
	totalBalance := totalCredits - totalDebits
	return successfulTxns, totalBalance, nil
}

func (t *TransactionService) GetAllIssues(ctx context.Context) ([]model.Transaction, error) {
	transactions, err := t.respository.GetTransaction().GetAllTransactions(ctx)
	if err != nil {
		return nil, err
	}
	var problemTxns []model.Transaction
	for _, txn := range transactions {
		if strings.ToLower(txn.Status) != "success" {
			problemTxns = append(problemTxns, txn)
		}
	}
	return problemTxns, nil
}

func (t *TransactionService) UploadTransactionCSV(ctx context.Context, reader io.Reader) (error, []string) {
	csvReader := csv.NewReader(reader)
	records, err := csvReader.ReadAll()
	if err != nil {
		return fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrInvalidCSVFormat), err), nil
	}

	if len(records) == 0 {
		return fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrEmptyFile), "Content is empty"), nil
	}

	var (
		txns         []model.Transaction
		errorDetails []string
	)

	for i, record := range records {
		//Check for missing/empty fields
		missing := false
		for j, field := range record {
			if strings.TrimSpace(field) == "" {
				errorDetails = append(errorDetails, fmt.Sprintf("Line %d: missing field at position %d", i+1, j+1))
				missing = true
				break
			}
		}
		if missing {
			continue
		}

		//Parse the record
		txn, err := parseTransactionRecord(record)
		if err != nil {
			errorDetails = append(errorDetails, fmt.Sprintf("Line %d: %v", i+1, err))
			continue
		}

		txns = append(txns, txn)
	}

	if len(txns) == 0 {
		msg := "No valid transactions found."
		if len(errorDetails) > 0 {
			msg += " Errors: " + strings.Join(errorDetails, "; ")
		}
		return fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrEmptyFile), msg), errorDetails
	}

	t.respository.GetTransaction().SaveTransactions(ctx, txns)
	return nil, errorDetails
}

func parseTransactionRecord(record []string) (model.Transaction, error) {
	timestamp, err := strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrInvalidFieldFormat), err)
	}

	amount, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return model.Transaction{}, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrInvalidFieldFormat), err)
	}

	return model.Transaction{
		Timestamp:   timestamp,
		Name:        record[1],
		Type:        record[2],
		Amount:      amount,
		Status:      record[4],
		Description: record[5],
	}, nil
}
