package controllers

import (
	transactionControllers "go-transaction-service/controllers/transaction"
	"go-transaction-service/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetTransaction() transactionControllers.ITransactionController
}

func NewControllerRegistry(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetTransaction() transactionControllers.ITransactionController {
	return transactionControllers.NewTransactionController(r.service)
}
