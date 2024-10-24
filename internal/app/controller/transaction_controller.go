package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type TransactionController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type TransactionControllerImpl struct {
	transactionService service.TransactionService
}

func (controller TransactionControllerImpl) GetAll(c *gin.Context) {
	controller.transactionService.GetAll(c)
}

func (controller TransactionControllerImpl) GetByID(c *gin.Context) {
	controller.transactionService.GetByID(c)
}

func (controller TransactionControllerImpl) Create(c *gin.Context) {
	controller.transactionService.Create(c)
}

func (controller TransactionControllerImpl) Update(c *gin.Context) {
	controller.transactionService.Update(c)
}

func (controller TransactionControllerImpl) Delete(c *gin.Context) {
	controller.transactionService.Delete(c)
}

func TransactionControllerInit(transactionService service.TransactionService) *TransactionControllerImpl {
	return &TransactionControllerImpl{transactionService: transactionService}
}
