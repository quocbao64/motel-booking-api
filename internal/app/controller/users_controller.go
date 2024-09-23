package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetByEmail(c *gin.Context)
	GetByPhone(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UserControllerImpl struct {
	customerSvc service.UserService
}

func (cus UserControllerImpl) GetAll(c *gin.Context) {
	cus.customerSvc.GetAll(c)
}

func (cus UserControllerImpl) GetByID(c *gin.Context) {
	cus.customerSvc.GetByID(c)
}

func (cus UserControllerImpl) GetByEmail(c *gin.Context) {
	cus.customerSvc.GetByEmail(c)
}

func (cus UserControllerImpl) Create(c *gin.Context) {
	cus.customerSvc.Create(c)
}

func (cus UserControllerImpl) Update(c *gin.Context) {
	cus.customerSvc.Update(c)
}

func (cus UserControllerImpl) Delete(c *gin.Context) {
	cus.customerSvc.Delete(c)
}

func (cus UserControllerImpl) GetByPhone(c *gin.Context) {
	cus.customerSvc.GetByPhone(c)
}

func UserControllerInit(UserService service.UserService) *UserControllerImpl {
	return &UserControllerImpl{
		customerSvc: UserService,
	}
}
