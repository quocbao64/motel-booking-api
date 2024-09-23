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

type RoomService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type RoomServiceImpl struct {
	roomRepo    repository.RoomRepository
	addressRepo repository.AddressRepository
}

type RoomParams struct {
	Title   string `json:"title" form:"title"`
	PageID  int    `json:"page_id" form:"page_id" binding:"required"`
	PerPage int    `json:"per_page" form:"per_page" binding:"required"`
}

func (repo RoomServiceImpl) GetAll(c *gin.Context) {

	var params RoomParams
	err := c.BindQuery(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	filter := &repository.RoomFilter{
		Title:   params.Title,
		PageID:  params.PageID,
		PerPage: params.PerPage,
	}
	data, err := repo.roomRepo.GetAll(filter)

	for _, room := range data {
		address, err := repo.addressRepo.GetFullAddress(room.AddressID)
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}
		room.Address = *address
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo RoomServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.roomRepo.GetByID(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo RoomServiceImpl) Create(c *gin.Context) {
	var room *dao.Room
	err := c.BindJSON(&room)

	if err != nil {
		return
	}

	data, err := repo.roomRepo.Create(room)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo RoomServiceImpl) Update(c *gin.Context) {
	var room *dao.Room
	err := c.BindJSON(&room)

	if err != nil {
		return
	}

	data, err := repo.roomRepo.Update(room)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo RoomServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.roomRepo.Delete(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func RoomServiceInit(roomRepo repository.RoomRepository, addressRepo repository.AddressRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		roomRepo:    roomRepo,
		addressRepo: addressRepo,
	}
}
