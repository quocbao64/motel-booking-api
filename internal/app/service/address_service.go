package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddressService interface {
	GetFullAddress(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
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

func (service AddressServiceImpl) Create(c *gin.Context) {
	var address *dao.Address
	err := c.ShouldBindJSON(&address)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	data, err := service.geographyRepo.Create(address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (service AddressServiceImpl) Update(c *gin.Context) {
	var address *dao.Address
	var id, _ = strconv.Atoi(c.Param("id"))
	err := c.ShouldBindJSON(&address)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	address.ID = uint(id)

	data, err := service.geographyRepo.Update(address)

	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func AddressServiceInit(geographyRepo repository.AddressRepository) *AddressServiceImpl {
	return &AddressServiceImpl{
		geographyRepo: geographyRepo,
	}
}
