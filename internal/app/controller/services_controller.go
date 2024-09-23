package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type ServicesController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ServicesControllerImpl struct {
	servicesService service.ServicesService
}

func (controller ServicesControllerImpl) GetAll(c *gin.Context) {
	controller.servicesService.GetAll(c)
}

func (controller ServicesControllerImpl) GetByID(c *gin.Context) {
	controller.servicesService.GetByID(c)
}

func (controller ServicesControllerImpl) Create(c *gin.Context) {
	controller.servicesService.Create(c)
}

func (controller ServicesControllerImpl) Update(c *gin.Context) {
	controller.servicesService.Update(c)
}

func (controller ServicesControllerImpl) Delete(c *gin.Context) {
	controller.servicesService.Delete(c)
}

func ServicesControllerInit(servicesService service.ServicesService) *ServicesControllerImpl {
	return &ServicesControllerImpl{servicesService: servicesService}
}
