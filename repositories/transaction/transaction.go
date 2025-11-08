package repositories

import (
	"context"
	"fmt"
	errWrap "go-transaction-service/common/error"
	errConstant "go-transaction-service/constants/error"
	"go-transaction-service/domain/model"
)

type ITransactionRepository interface {
	SaveTransactions(ctx context.Context, txns []model.Transaction) []model.Transaction
	GetAllTransactions(ctx context.Context) ([]model.Transaction, error)
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

func (t *TransactionRepository) GetAllTransactions(ctx context.Context) ([]model.Transaction, error) {
	if len(t.data) == 0 {
		return nil, fmt.Errorf("%w: %v", errWrap.WrapError(errConstant.ErrEmptyData), "No transaction data available.")
	}
	return t.data, nil
}
