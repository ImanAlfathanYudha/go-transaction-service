package routes

import (
	"go-transaction-service/controllers"
	transactionRoutes "go-transaction-service/routes/transaction"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type IRouterRegister interface {
	Serve()
}

func NewRouterRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRouterRegister {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) Serve() {
	r.transationRoute().Run()
}

func (r *Registry) transationRoute() transactionRoutes.ITransactionRoute {
	return transactionRoutes.NewTransactionRoute(r.controller, r.group)
}
