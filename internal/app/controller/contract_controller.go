package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type ContractController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ContractControllerImpl struct {
	contractService service.ContractService
}

func (controller ContractControllerImpl) GetAll(c *gin.Context) {
	controller.contractService.GetAll(c)
}

func (controller ContractControllerImpl) GetByID(c *gin.Context) {
	controller.contractService.GetByID(c)
}

func (controller ContractControllerImpl) Create(c *gin.Context) {
	controller.contractService.Create(c)
}

func (controller ContractControllerImpl) Update(c *gin.Context) {
	controller.contractService.Update(c)
}

func (controller ContractControllerImpl) Delete(c *gin.Context) {
	controller.contractService.Delete(c)
}

func ContractControllerInit(contractService service.ContractService) *ContractControllerImpl {
	return &ContractControllerImpl{contractService: contractService}
}
