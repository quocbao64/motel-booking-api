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

type ServicesService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ServicesServiceImpl struct {
	servicesRepo repository.ServiceRepository
}

func (repo ServicesServiceImpl) GetAll(c *gin.Context) {
	data, err := repo.servicesRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.servicesRepo.GetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesServiceImpl) Create(c *gin.Context) {
	var service *dao.Service
	err := c.BindJSON(&service)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.servicesRepo.Create(service)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesServiceImpl) Update(c *gin.Context) {
	var service *dao.Service
	err := c.BindJSON(&service)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.servicesRepo.Update(service)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.servicesRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func ServicesServiceInit(repo repository.ServiceRepository) *ServicesServiceImpl {
	return &ServicesServiceImpl{servicesRepo: repo}
}
