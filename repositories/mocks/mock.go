// in go-transaction-service/repositories/mocks/transaction_repository_mock.go

package mocks

import (
	"context"
	"go-transaction-service/domain/model"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (m *TransactionRepositoryMock) SaveTransactions(ctx context.Context, txns []model.Transaction) []model.Transaction {
	args := m.Called(ctx, txns)
	return args.Get(0).([]model.Transaction)
}

func (m *TransactionRepositoryMock) GetAllTransactions(ctx context.Context) ([]model.Transaction, error) {
	args := m.Called(ctx)
	return args.Get(0).([]model.Transaction), args.Error(1)
}
