package routes

import (
	"awesomeProject/config"
	_ "awesomeProject/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func Route(init *config.Initialize) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3001"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		MaxAge:       12 * time.Hour,
	}))

	api := router.Group("/api/v1")
	{
		UserRoute(init, api.Group("/users"))
		AuthRoute(init, api.Group("/auth"))
		RoomRoute(init, api.Group("/rooms"))
		GeographyRoute(init, api.Group("/geography"))
		AddressRoute(init, api.Group("/address"))
		ServiceRoute(init, api.Group("/services"))
		ContractRoute(init, api.Group("/contracts"))
		InvoiceRoute(init, api.Group("/invoices"))
		ServiceDemandRoute(init, api.Group("/services-demand"))
		BookingRequestRoute(init, api.Group("/booking-requests"))
	}

	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
