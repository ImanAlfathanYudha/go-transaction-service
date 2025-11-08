package services_test

import (
	"context"
	"strings"
	"testing"

	errWrap "go-transaction-service/common/error"
	errConstant "go-transaction-service/constants/error"
	"go-transaction-service/domain/model"
	"go-transaction-service/repositories/mocks"
	repositories "go-transaction-service/repositories/transaction"
	"go-transaction-service/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepositoryRegistry struct {
	transactionRepo repositories.ITransactionRepository
}

func (m *MockRepositoryRegistry) GetTransaction() repositories.ITransactionRepository {
	return m.transactionRepo
}

func TestUploadTransactionCSV_Success(t *testing.T) {
	mockRepo := new(mocks.TransactionRepositoryMock)
	mockRegistry := &MockRepositoryRegistry{transactionRepo: mockRepo}
	service := services.NewServiceRegistry(mockRegistry)

	csvData := `1624512883,COMPANY A,CREDIT,12000000,SUCCESS,salary
1624700000,COFFEE SHOP,DEBIT,50000,SUCCESS,coffee`

	reader := strings.NewReader(csvData)
	ctx := context.Background()

	mockRepo.On("SaveTransactions", ctx, mock.Anything).Return([]model.Transaction{})

	err, _ := service.GetTransaction().UploadTransactionCSV(ctx, reader)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "SaveTransactions", ctx, mock.Anything)
}

func TestUploadTransactionCSV_EmptyFile(t *testing.T) {
	mockRepo := new(mocks.TransactionRepositoryMock)
	mockRegistry := &MockRepositoryRegistry{transactionRepo: mockRepo}
	service := services.NewServiceRegistry(mockRegistry)

	reader := strings.NewReader("") // empty CSV
	ctx := context.Background()

	err, _ := service.GetTransaction().UploadTransactionCSV(ctx, reader)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), errWrap.WrapError(errConstant.ErrEmptyFile).Error())
}

func TestUploadTransactionCSV_InvalidData(t *testing.T) {
	mockRepo := new(mocks.TransactionRepositoryMock)
	mockRegistry := &MockRepositoryRegistry{transactionRepo: mockRepo}
	service := services.NewServiceRegistry(mockRegistry)

	// Invalid data: amount is string instead of number
	csvData := `1624512883,COMPANY A,CREDIT,abc,SUCCESS,salary`
	reader := strings.NewReader(csvData)
	ctx := context.Background()

	err, _ := service.GetTransaction().UploadTransactionCSV(ctx, reader)
	assert.Error(t, err)
}
