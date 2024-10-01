package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type BookingRequestController interface {
	GetAllBookingRequest(c *gin.Context)
	GetBookingRequestByID(c *gin.Context)
	CreateBookingRequest(c *gin.Context)
	UpdateBookingRequest(c *gin.Context)
	DeleteBookingRequest(c *gin.Context)
}

type BookingRequestControllerImpl struct {
	bookingRequestService service.BookingRequestService
}

func (controller BookingRequestControllerImpl) GetAllBookingRequest(c *gin.Context) {
	controller.bookingRequestService.GetAll(c)
}

func (controller BookingRequestControllerImpl) GetBookingRequestByID(c *gin.Context) {
	controller.bookingRequestService.GetByID(c)
}

func (controller BookingRequestControllerImpl) CreateBookingRequest(c *gin.Context) {
	controller.bookingRequestService.Create(c)
}

func (controller BookingRequestControllerImpl) UpdateBookingRequest(c *gin.Context) {
	controller.bookingRequestService.Update(c)
}

func (controller BookingRequestControllerImpl) DeleteBookingRequest(c *gin.Context) {
	controller.bookingRequestService.Delete(c)
}

func BookingRequestControllerInit(bookingRequestService service.BookingRequestService) *BookingRequestControllerImpl {
	return &BookingRequestControllerImpl{bookingRequestService: bookingRequestService}
}
