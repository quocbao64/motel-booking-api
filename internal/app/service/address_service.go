package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddressService interface {
	GetFullAddress(c *gin.Context)
}

type AddressServiceImpl struct {
	geographyRepo repository.AddressRepository
}

func (service AddressServiceImpl) GetFullAddress(c *gin.Context) {
	id := c.Param("id")

	uintID, _ := strconv.Atoi(id)

	data, err := service.geographyRepo.GetFullAddress(uint(uintID))

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func AddressServiceInit(geographyRepo repository.AddressRepository) *AddressServiceImpl {
	return &AddressServiceImpl{
		geographyRepo: geographyRepo,
	}
}
