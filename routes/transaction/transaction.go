package transaction

import (
	"go-transaction-service/controllers"

	"github.com/gin-gonic/gin"
)

type TransactionRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
}

type ITransactionRoute interface {
	Run()
}

func NewTransactionRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup) ITransactionRoute {
	return &TransactionRoute{controller: controller, group: group}
}

func (u *TransactionRoute) Run() {
	group := u.group.Group("/transaction")
	group.POST("/upload", u.controller.GetTransaction().UploadTransactionCSV)
	group.GET("/balance", u.controller.GetTransaction().GetAllBalance)
	group.GET("/issues", u.controller.GetTransaction().GetAllIssues)
}
