package config

import (
	"awesomeProject/internal/app/controller"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/service"
)

type Initialize struct {
	AuthSvc            service.AuthService
	AuthCtrl           controller.AuthController
	UserCtrl           controller.UserController
	UserSvc            service.UserService
	UserRepo           repository.UserRepository
	RoomRepo           repository.RoomRepository
	RoomCtrl           controller.RoomController
	RoomSvc            service.RoomService
	GeographyRepo      repository.GeographyRepository
	GeographyCtrl      controller.GeographyController
	GeographySvc       service.GeographyService
	AddressCtrl        controller.AddressController
	AddressSvc         service.AddressService
	AddressRepo        repository.AddressRepository
	ServiceRepo        repository.ServiceRepository
	ServiceCtrl        controller.ServicesController
	ServiceSvc         service.ServicesService
	HashContractSvc    service.HashContractService
	HashContractRepo   repository.HashContractRepository
	ContractCtrl       controller.ContractController
	ContractSvc        service.ContractService
	ContractRepo       repository.ContractRepository
	InvoiceCtrl        controller.InvoiceController
	InvoiceSvc         service.InvoiceService
	InvoiceRepo        repository.InvoiceRepository
	ServicesDemandCtrl controller.ServicesDemandController
	ServicesDemandSvc  service.ServicesDemandService
	ServicesDemandRepo repository.ServicesDemandRepository
	BookingRequestCtrl controller.BookingRequestController
	BookingRequestSvc  service.BookingRequestService
	BookingRequestRepo repository.BookingRequestRepository
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
	hashContractSvc service.HashContractService,
	hashContractRepo repository.HashContractRepository,
	contractCtrl controller.ContractController,
	contractSvc service.ContractService,
	contractRepo repository.ContractRepository,
	invoiceCtrl controller.InvoiceController,
	invoiceSvc service.InvoiceService,
	invoiceRepo repository.InvoiceRepository,
	servicesDemandCtrl controller.ServicesDemandController,
	servicesDemandSvc service.ServicesDemandService,
	servicesDemandRepo repository.ServicesDemandRepository,
	bookingRequestCtrl controller.BookingRequestController,
	bookingRequestSvc service.BookingRequestService,
	bookingRequestRepo repository.BookingRequestRepository,
) *Initialize {
	return &Initialize{
		AuthCtrl:           authCtrl,
		AuthSvc:            authSvc,
		UserCtrl:           userCtrl,
		UserSvc:            userSvc,
		UserRepo:           userRepo,
		RoomCtrl:           roomCtrl,
		RoomSvc:            roomSvc,
		RoomRepo:           roomRepo,
		GeographyRepo:      geographyRepo,
		GeographyCtrl:      geographyCtrl,
		GeographySvc:       geographySvc,
		AddressCtrl:        addressCtrl,
		AddressSvc:         addressSvc,
		AddressRepo:        addressRepo,
		ServiceRepo:        serviceRepo,
		ServiceCtrl:        serviceCtrl,
		ServiceSvc:         serviceSvc,
		HashContractSvc:    hashContractSvc,
		HashContractRepo:   hashContractRepo,
		ContractCtrl:       contractCtrl,
		ContractSvc:        contractSvc,
		ContractRepo:       contractRepo,
		InvoiceCtrl:        invoiceCtrl,
		InvoiceSvc:         invoiceSvc,
		InvoiceRepo:        invoiceRepo,
		ServicesDemandCtrl: servicesDemandCtrl,
		ServicesDemandSvc:  servicesDemandSvc,
		ServicesDemandRepo: servicesDemandRepo,
		BookingRequestCtrl: bookingRequestCtrl,
		BookingRequestSvc:  bookingRequestSvc,
		BookingRequestRepo: bookingRequestRepo,
	}
}
