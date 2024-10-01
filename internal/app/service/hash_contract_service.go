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

type HashContractService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type HashContractServiceImpl struct {
	contractRepo repository.HashContractRepository
}

func (repo HashContractServiceImpl) GetAll(c *gin.Context) {
	data, err := repo.contractRepo.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo HashContractServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.contractRepo.GetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo HashContractServiceImpl) Create(c *gin.Context) {
	var contract *dao.HashContract
	err := c.BindJSON(&contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.contractRepo.Create(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo HashContractServiceImpl) Update(c *gin.Context) {
	var contract *dao.HashContract
	err := c.BindJSON(&contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.contractRepo.Update(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo HashContractServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.contractRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func HashContractServiceInit(repo repository.HashContractRepository) *HashContractServiceImpl {
	return &HashContractServiceImpl{contractRepo: repo}
}
