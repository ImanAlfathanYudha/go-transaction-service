package services

import (
	transactionRepo "go-transaction-service/repositories"
	transactionService "go-transaction-service/services/transaction"
)

type Registry struct {
	repository transactionRepo.IRepositoryRegistry
}

type IServiceRegistry interface {
	GetTransaction() transactionService.ITransactionService
}

func NewServiceRegistry(repository transactionRepo.IRepositoryRegistry) IServiceRegistry {
	return &Registry{repository: repository}
}

func (r *Registry) GetTransaction() transactionService.ITransactionService {
	return transactionService.NewTransactionService(r.repository)
}
