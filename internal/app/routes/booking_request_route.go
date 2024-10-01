package routes

import (
	"awesomeProject/config"
	"github.com/gin-gonic/gin"
)

func BookingRequestRoute(init *config.Initialize, g *gin.RouterGroup) *gin.RouterGroup {
	g.GET("", init.BookingRequestCtrl.GetAllBookingRequest)
	g.POST("", init.BookingRequestCtrl.CreateBookingRequest)
	g.GET("/:id", init.BookingRequestCtrl.GetBookingRequestByID)
	g.PUT("/:id", init.BookingRequestCtrl.UpdateBookingRequest)
	g.DELETE("/:id", init.BookingRequestCtrl.DeleteBookingRequest)
	return g
}
