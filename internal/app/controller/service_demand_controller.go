package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type ServicesDemandController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ServicesDemandControllerImpl struct {
	servicesDemandService service.ServicesDemandService
}

func (controller ServicesDemandControllerImpl) GetAll(c *gin.Context) {
	controller.servicesDemandService.GetAll(c)
}

func (controller ServicesDemandControllerImpl) GetByID(c *gin.Context) {
	controller.servicesDemandService.GetByID(c)
}

func (controller ServicesDemandControllerImpl) Create(c *gin.Context) {
	controller.servicesDemandService.Create(c)
}

func (controller ServicesDemandControllerImpl) Update(c *gin.Context) {
	controller.servicesDemandService.Update(c)
}

func (controller ServicesDemandControllerImpl) Delete(c *gin.Context) {
	controller.servicesDemandService.Delete(c)
}

func ServicesDemandControllerInit(servicesDemandService service.ServicesDemandService) *ServicesDemandControllerImpl {
	return &ServicesDemandControllerImpl{servicesDemandService: servicesDemandService}
}
