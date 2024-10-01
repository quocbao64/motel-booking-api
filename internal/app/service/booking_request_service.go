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

type BookingRequestService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type BookingRequestServiceImpl struct {
	bookingRequestRepo repository.BookingRequestRepository
	roomRepo           repository.RoomRepository
}

type BookingRequestParams struct {
	RenterID int `json:"renter_id" form:"renter_id"`
	LessorID int `json:"lessor_id" form:"lessor_id"`
	PageID   int `json:"page_id" form:"page_id"`
	PerPage  int `json:"per_page" form:"per_page"`
}

func (repo BookingRequestServiceImpl) GetAll(c *gin.Context) {
	params := &BookingRequestParams{}
	err := c.Bind(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	filter := &repository.BookingRequestFilter{
		RenterID: params.RenterID,
		LessorID: params.LessorID,
	}

	data, err := repo.bookingRequestRepo.GetAll(filter)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo BookingRequestServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.bookingRequestRepo.GetByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo BookingRequestServiceImpl) Create(c *gin.Context) {
	var bookingRequest *dao.BookingRequest
	err := c.BindJSON(&bookingRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.bookingRequestRepo.Create(bookingRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo BookingRequestServiceImpl) Update(c *gin.Context) {
	var bookingRequest *dao.BookingRequest
	err := c.BindJSON(&bookingRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.bookingRequestRepo.Update(bookingRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo BookingRequestServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.bookingRequestRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func (repo BookingRequestServiceImpl) GetByRenterOrLessorID(c *gin.Context) {
	renterID, _ := strconv.Atoi(c.Query("renter_id"))
	lessorID, _ := strconv.Atoi(c.Query("lessor_id"))
	data, err := repo.bookingRequestRepo.GetByRenterOrLessorID(renterID, lessorID)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func BookingRequestServiceInit(bookingRequestRepo repository.BookingRequestRepository) *BookingRequestServiceImpl {
	return &BookingRequestServiceImpl{bookingRequestRepo: bookingRequestRepo}
}