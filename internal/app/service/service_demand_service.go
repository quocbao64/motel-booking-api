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

type ServicesDemandService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ServicesDemandServiceImpl struct {
	serviceDemandRepo repository.ServicesDemandRepository
}

func (repo ServicesDemandServiceImpl) GetAll(c *gin.Context) {
	data, err := repo.serviceDemandRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesDemandServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.serviceDemandRepo.GetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesDemandServiceImpl) Create(c *gin.Context) {
	var serviceDemand *dao.ServiceDemand
	err := c.BindJSON(&serviceDemand)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.serviceDemandRepo.Create(serviceDemand)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesDemandServiceImpl) Update(c *gin.Context) {
	var serviceDemand *dao.ServiceDemand
	err := c.BindJSON(&serviceDemand)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.serviceDemandRepo.Update(serviceDemand)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ServicesDemandServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.serviceDemandRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func ServicesDemandServiceInit(repo repository.ServicesDemandRepository) *ServicesDemandServiceImpl {
	return &ServicesDemandServiceImpl{serviceDemandRepo: repo}
}
