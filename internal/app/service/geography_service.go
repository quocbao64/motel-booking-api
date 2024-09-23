package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type GeographyService interface {
	GetAllDistrict(c *gin.Context)
	GetDistrictByID(c *gin.Context)
	GetAllProvince(c *gin.Context)
	GetProvinceByID(c *gin.Context)
	GetAllWard(c *gin.Context)
	GetWardByID(c *gin.Context)
}

type GeographyServiceImpl struct {
	geographyRepo repository.GeographyRepository
}

func (repo GeographyServiceImpl) GetAllDistrict(c *gin.Context) {
	data, err := repo.geographyRepo.GetAllDistrict()

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo GeographyServiceImpl) GetDistrictByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.geographyRepo.GetDistrictByID(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo GeographyServiceImpl) GetAllProvince(c *gin.Context) {
	data, err := repo.geographyRepo.GetAllProvince()

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo GeographyServiceImpl) GetProvinceByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.geographyRepo.GetProvinceByID(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo GeographyServiceImpl) GetAllWard(c *gin.Context) {
	data, err := repo.geographyRepo.GetAllWard()

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo GeographyServiceImpl) GetWardByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.geographyRepo.GetWardByID(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func GeographyServiceInit(geographyRepo repository.GeographyRepository) *GeographyServiceImpl {
	return &GeographyServiceImpl{
		geographyRepo: geographyRepo,
	}
}
