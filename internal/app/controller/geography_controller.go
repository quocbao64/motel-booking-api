package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type GeographyController interface {
	GetAllDistrict(c *gin.Context)
	GetDistrictByID(c *gin.Context)
	GetAllProvince(c *gin.Context)
	GetProvinceByID(c *gin.Context)
	GetAllWard(c *gin.Context)
	GetWardByID(c *gin.Context)
}

type GeographyControllerImpl struct {
	geographySvc service.GeographyService
}

func (cus GeographyControllerImpl) GetAllDistrict(c *gin.Context) {
	cus.geographySvc.GetAllDistrict(c)
}

func (cus GeographyControllerImpl) GetDistrictByID(c *gin.Context) {
	cus.geographySvc.GetDistrictByID(c)
}

func (cus GeographyControllerImpl) GetAllProvince(c *gin.Context) {
	cus.geographySvc.GetAllProvince(c)
}

func (cus GeographyControllerImpl) GetProvinceByID(c *gin.Context) {
	cus.geographySvc.GetProvinceByID(c)
}

func (cus GeographyControllerImpl) GetAllWard(c *gin.Context) {
	cus.geographySvc.GetAllWard(c)
}

func (cus GeographyControllerImpl) GetWardByID(c *gin.Context) {
	cus.geographySvc.GetWardByID(c)
}

func GeographyControllerInit(GeographyService service.GeographyService) *GeographyControllerImpl {
	return &GeographyControllerImpl{
		geographySvc: GeographyService,
	}
}
