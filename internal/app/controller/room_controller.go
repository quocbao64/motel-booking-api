package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type RoomController interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type RoomControllerImpl struct {
	roomSvc service.RoomService
}

func (cus RoomControllerImpl) GetAll(c *gin.Context) {
	cus.roomSvc.GetAll(c)
}

func (cus RoomControllerImpl) GetByID(c *gin.Context) {
	cus.roomSvc.GetByID(c)
}

func (cus RoomControllerImpl) Create(c *gin.Context) {
	cus.roomSvc.Create(c)
}

func (cus RoomControllerImpl) Update(c *gin.Context) {
	cus.roomSvc.Update(c)
}

func (cus RoomControllerImpl) Delete(c *gin.Context) {
	cus.roomSvc.Delete(c)
}

func RoomControllerInit(RoomService service.RoomService) *RoomControllerImpl {
	return &RoomControllerImpl{
		roomSvc: RoomService,
	}
}
