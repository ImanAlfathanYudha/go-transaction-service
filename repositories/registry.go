package repositories

import (
	transactionRepo "go-transaction-service/repositories/transaction"
)

type IRepositoryRegistry interface {
	GetTransaction() transactionRepo.ITransactionRepository
}

type Registry struct {
	transactionRepo transactionRepo.ITransactionRepository
}

func NewRepositoryRegistry() IRepositoryRegistry {
	return &Registry{
		transactionRepo: transactionRepo.NewTransactionRepository(),
	}
}

func (r *Registry) GetTransaction() transactionRepo.ITransactionRepository {
	return r.transactionRepo
}
