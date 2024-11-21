package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type RoomService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	UpdateStatus(c *gin.Context)
}

type RoomServiceImpl struct {
	roomRepo         repository.RoomRepository
	addressRepo      repository.AddressRepository
	borrowedItemRepo repository.BorrowedItemRepository
	serviceRepo      repository.ServiceRepository
}

type RoomParams struct {
	Title   string `json:"title" form:"title"`
	PageID  int    `json:"page_id" form:"page_id" binding:"required"`
	PerPage int    `json:"per_page" form:"per_page" binding:"required"`
}

type RoomRequest struct {
	ID     int `json:"id"`
	Status int `json:"status"`
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

	address, err := repo.addressRepo.GetFullAddress(data.AddressID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}
	data.Address = *address

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo RoomServiceImpl) Create(c *gin.Context) {
	var roomReq *dao.RoomRequest
	err := c.BindJSON(&roomReq)

	if err != nil {
		return
	}

	var images []string
	if roomReq.Images != nil {
		for _, image := range roomReq.Images {
			url, err := pkg.UploadS3("rooms/"+uuid.New().String()+"/"+image.FileName, []byte(image.FileBase64), "image")
			if err != nil {
				return
			}
			images = append(images, url)
		}
	}

	address, err := repo.addressRepo.Create(&dao.Address{
		WardID: roomReq.WardID,
		Detail: roomReq.AddressDetail,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	room := &dao.Room{
		Title:         roomReq.Title,
		Acreage:       roomReq.Acreage,
		Price:         roomReq.Price,
		Description:   roomReq.Description,
		DateSubmitted: roomReq.DateSubmitted,
		OwnerID:       roomReq.OwnerID,
		MaxPeople:     roomReq.MaxPeople,
		RoomType:      roomReq.RoomType,
		Deposit:       roomReq.Deposit,
		Images:        images,
		Address:       *address,
		Status:        roomReq.Status,
	}

	data, err := repo.roomRepo.Create(room)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	if roomReq.Services != nil {
		for _, service := range roomReq.Services {
			_, err := repo.serviceRepo.Create(&service)
			if err != nil {
				c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
				return
			}

			err = repo.roomRepo.CreateRoomService(data.ID, service.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
				return
			}
		}
	}

	if roomReq.BorrowedItems != nil {
		for _, borrowedItem := range roomReq.BorrowedItems {
			borrowedItem.RoomID = data.ID
			_, err := repo.borrowedItemRepo.Create(&borrowedItem)
			if err != nil {
				c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
				return
			}
		}
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

func (repo RoomServiceImpl) UpdateStatus(c *gin.Context) {
	var roomReq *RoomRequest
	err := c.BindJSON(&roomReq)

	if err != nil {
		return
	}

	err = repo.roomRepo.UpdateStatus(roomReq.ID, roomReq.Status)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func RoomServiceInit(roomRepo repository.RoomRepository,
	addressRepo repository.AddressRepository,
	borrowedItemRepo repository.BorrowedItemRepository,
	serviceRepo repository.ServiceRepository) *RoomServiceImpl {
	return &RoomServiceImpl{
		roomRepo:         roomRepo,
		addressRepo:      addressRepo,
		borrowedItemRepo: borrowedItemRepo,
		serviceRepo:      serviceRepo,
	}
}
