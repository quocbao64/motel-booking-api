package config

import (
	"awesomeProject/internal/app/controller"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/service"
)

type Initialize struct {
	AuthSvc       service.AuthService
	AuthCtrl      controller.AuthController
	UserCtrl      controller.UserController
	UserSvc       service.UserService
	UserRepo      repository.UserRepository
	RoomRepo      repository.RoomRepository
	RoomCtrl      controller.RoomController
	RoomSvc       service.RoomService
	GeographyRepo repository.GeographyRepository
	GeographyCtrl controller.GeographyController
	GeographySvc  service.GeographyService
	AddressCtrl   controller.AddressController
	AddressSvc    service.AddressService
	AddressRepo   repository.AddressRepository
	ServiceRepo   repository.ServiceRepository
	ServiceCtrl   controller.ServicesController
	ServiceSvc    service.ServicesService
}

func NewInitialize(
	authCtrl controller.AuthController,
	authSvc service.AuthService,
	userCtrl controller.UserController,
	userSvc service.UserService,
	userRepo repository.UserRepository,
	roomCtrl controller.RoomController,
	roomSvc service.RoomService,
	roomRepo repository.RoomRepository,
	geographyRepo repository.GeographyRepository,
	geographyCtrl controller.GeographyController,
	geographySvc service.GeographyService,
	addressCtrl controller.AddressController,
	addressSvc service.AddressService,
	addressRepo repository.AddressRepository,
	serviceRepo repository.ServiceRepository,
	serviceCtrl controller.ServicesController,
	serviceSvc service.ServicesService,
) *Initialize {
	return &Initialize{
		AuthCtrl:      authCtrl,
		AuthSvc:       authSvc,
		UserCtrl:      userCtrl,
		UserSvc:       userSvc,
		UserRepo:      userRepo,
		RoomCtrl:      roomCtrl,
		RoomSvc:       roomSvc,
		RoomRepo:      roomRepo,
		GeographyRepo: geographyRepo,
		GeographyCtrl: geographyCtrl,
		GeographySvc:  geographySvc,
		AddressCtrl:   addressCtrl,
		AddressSvc:    addressSvc,
		AddressRepo:   addressRepo,
		ServiceRepo:   serviceRepo,
		ServiceCtrl:   serviceCtrl,
		ServiceSvc:    serviceSvc,
	}
}
