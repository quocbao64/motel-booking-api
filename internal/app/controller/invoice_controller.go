package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type InvoiceController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type InvoiceControllerImpl struct {
	invoiceService service.InvoiceService
}

func (controller InvoiceControllerImpl) GetAll(c *gin.Context) {
	controller.invoiceService.GetAll(c)
}

func (controller InvoiceControllerImpl) GetByID(c *gin.Context) {
	controller.invoiceService.GetByID(c)
}

func (controller InvoiceControllerImpl) Create(c *gin.Context) {
	controller.invoiceService.Create(c)
}

func (controller InvoiceControllerImpl) Update(c *gin.Context) {
	controller.invoiceService.Update(c)
}

func (controller InvoiceControllerImpl) Delete(c *gin.Context) {
	controller.invoiceService.Delete(c)
}

func InvoiceControllerInit(invoiceService service.InvoiceService) *InvoiceControllerImpl {
	return &InvoiceControllerImpl{invoiceService: invoiceService}
}
