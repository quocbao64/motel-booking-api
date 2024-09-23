package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type AddressController interface {
	GetFullAddress(c *gin.Context)
}

type AddressControllerImpl struct {
	geographySvc service.AddressService
}

func (cus AddressControllerImpl) GetFullAddress(c *gin.Context) {
	cus.geographySvc.GetFullAddress(c)
}

func AddressControllerInit(AddressService service.AddressService) *AddressControllerImpl {
	return &AddressControllerImpl{
		geographySvc: AddressService,
	}
}
