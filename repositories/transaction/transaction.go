package repositories

import (
	"context"
	"fmt"
	errWrap "go-transaction-service/common/error"
	errConstant "go-transaction-service/constants/error"
	"go-transaction-service/domain/model"
	"strings"
)

type ITransactionRepository interface {
	SaveTransactions(ctx context.Context, txns []model.Transaction) []model.Transaction
	GetAllTransactions(ctx context.Context) ([]model.Transaction, error)
	EditTransactionStatus(ctx context.Context, timeStamp int64, status string) error
	GetTransactionByID(ctx context.Context, timeStamp int64) (model.Transaction, error)
	UpdateTransactionByID(ctx context.Context, timeStamp int64, request model.Transaction) error
}

type TransactionRepository struct {
	data []model.Transaction
}

func NewTransactionRepository() ITransactionRepository {
	return &TransactionRepository{}
}

func (t *TransactionRepository) SaveTransactions(ctx context.Context, txns []model.Transaction) []model.Transaction {
	t.data = append(t.data, txns...)
	return t.data
}

func (t *TransactionRepository) EditTransactionStatus(ctx context.Context, timeStamp int64, status string) error {
	for i := range t.data {
		if t.data[i].Timestamp == timeStamp {
			t.data[i].Status = strings.ToUpper(status)
			return nil
		}
	}
	return fmt.Errorf("transaction with timestamp %d not found", timeStamp)
}

func (t *TransactionRepository) GetAllTransactions(ctx context.Context) ([]model.Transaction, error) {
	if len(t.data) == 0 {
		return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrEmptyData), "No transaction data available.")
	}
	return t.data, nil
}

func (t *TransactionRepository) GetTransactionByID(ctx context.Context, timeStamp int64) (model.Transaction, error) {
	for i := range t.data {
		if t.data[i].Timestamp == timeStamp {
			return t.data[i], nil
		}
	}
	return model.Transaction{}, fmt.Errorf("transaction with timestamp %d not found", timeStamp)
}

func (t *TransactionRepository) UpdateTransactionByID(ctx context.Context, timeStamp int64, request model.Transaction) error {
	for i := range t.data {
		if t.data[i].Timestamp == timeStamp {
			t.data[i] = request
			return nil
		}
	}
	return fmt.Errorf("transaction with timestamp %d not found", timeStamp)
}
