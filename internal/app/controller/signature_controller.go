package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type SignatureController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type SignatureControllerImpl struct {
	transactionService service.SignatureService
}

func (controller SignatureControllerImpl) GetAll(c *gin.Context) {
	controller.transactionService.GetAll(c)
}

func (controller SignatureControllerImpl) GetByID(c *gin.Context) {
	controller.transactionService.GetByID(c)
}

func (controller SignatureControllerImpl) Create(c *gin.Context) {
	controller.transactionService.Create(c)
}

func (controller SignatureControllerImpl) Update(c *gin.Context) {
	controller.transactionService.Update(c)
}

func (controller SignatureControllerImpl) Delete(c *gin.Context) {
	controller.transactionService.Delete(c)
}

func SignatureControllerInit(transactionService service.SignatureService) *SignatureControllerImpl {
	return &SignatureControllerImpl{transactionService: transactionService}
}
