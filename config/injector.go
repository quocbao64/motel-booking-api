//go:build wireinject
// +build wireinject

package config

import (
	"awesomeProject/internal/app/controller"
	"awesomeProject/internal/app/repository"
	"awesomeProject/internal/app/service"
	"github.com/google/wire"
)

var db = wire.NewSet(ConnectDB)

var authSvcSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)))

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))

var addressSvcSet = wire.NewSet(service.AddressServiceInit,
	wire.Bind(new(service.AddressService), new(*service.AddressServiceImpl)))

var addressCtrlSet = wire.NewSet(controller.AddressControllerInit,
	wire.Bind(new(controller.AddressController), new(*controller.AddressControllerImpl)))

var addressRepoSet = wire.NewSet(repository.AddressRepositoryInit,
	wire.Bind(new(repository.AddressRepository), new(*repository.AddressRepositoryImpl)))

var userSvcSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var roomSvcSet = wire.NewSet(service.RoomServiceInit,
	wire.Bind(new(service.RoomService), new(*service.RoomServiceImpl)))

var roomCtrlSet = wire.NewSet(controller.RoomControllerInit,
	wire.Bind(new(controller.RoomController), new(*controller.RoomControllerImpl)))

var roomRepoSet = wire.NewSet(repository.RoomRepositoryInit,
	wire.Bind(new(repository.RoomRepository), new(*repository.RoomRepositoryImpl)))

var geographySvcSet = wire.NewSet(service.GeographyServiceInit,
	wire.Bind(new(service.GeographyService), new(*service.GeographyServiceImpl)))

var geographyCtrlSet = wire.NewSet(controller.GeographyControllerInit,
	wire.Bind(new(controller.GeographyController), new(*controller.GeographyControllerImpl)))

var geographyRepoSet = wire.NewSet(repository.GeographyRepositoryInit,
	wire.Bind(new(repository.GeographyRepository), new(*repository.GeographyRepositoryImpl)))

var serviceSvcSet = wire.NewSet(service.ServicesServiceInit,
	wire.Bind(new(service.ServicesService), new(*service.ServicesServiceImpl)))

var serviceCtrlSet = wire.NewSet(controller.ServicesControllerInit,
	wire.Bind(new(controller.ServicesController), new(*controller.ServicesControllerImpl)))

var serviceRepoSet = wire.NewSet(repository.ServiceRepositoryInit,
	wire.Bind(new(repository.ServiceRepository), new(*repository.ServiceRepositoryImpl)))

var contractSvcSet = wire.NewSet(service.ContractServiceInit,
	wire.Bind(new(service.ContractService), new(*service.ContractServiceImpl)))

var contractCtrlSet = wire.NewSet(controller.ContractControllerInit,
	wire.Bind(new(controller.ContractController), new(*controller.ContractControllerImpl)))

var contractRepoSet = wire.NewSet(repository.ContractRepositoryInit,
	wire.Bind(new(repository.ContractRepository), new(*repository.ContractRepositoryImpl)))

var invoiceSvcSet = wire.NewSet(service.InvoiceServiceInit,
	wire.Bind(new(service.InvoiceService), new(*service.InvoiceServiceImpl)))

var invoiceCtrlSet = wire.NewSet(controller.InvoiceControllerInit,
	wire.Bind(new(controller.InvoiceController), new(*controller.InvoiceControllerImpl)))

var invoiceRepoSet = wire.NewSet(repository.InvoiceRepositoryInit,
	wire.Bind(new(repository.InvoiceRepository), new(*repository.InvoiceRepositoryImpl)))

var hashContractRepoSet = wire.NewSet(repository.HashContractRepositoryInit,
	wire.Bind(new(repository.HashContractRepository), new(*repository.HashContractRepositoryImpl)))

var hashContractSvcSet = wire.NewSet(service.HashContractServiceInit,
	wire.Bind(new(service.HashContractService), new(*service.HashContractServiceImpl)))

var servicesDemandSvcSet = wire.NewSet(service.ServicesDemandServiceInit,
	wire.Bind(new(service.ServicesDemandService), new(*service.ServicesDemandServiceImpl)))

var servicesDemandCtrlSet = wire.NewSet(controller.ServicesDemandControllerInit,
	wire.Bind(new(controller.ServicesDemandController), new(*controller.ServicesDemandControllerImpl)))

var servicesDemandRepoSet = wire.NewSet(repository.ServicesDemandRepositoryInit,
	wire.Bind(new(repository.ServicesDemandRepository), new(*repository.ServicesDemandRepositoryImpl)))

var bookingRequestSvcSet = wire.NewSet(service.BookingRequestServiceInit,
	wire.Bind(new(service.BookingRequestService), new(*service.BookingRequestServiceImpl)))

var bookingRequestCtrlSet = wire.NewSet(controller.BookingRequestControllerInit,
	wire.Bind(new(controller.BookingRequestController), new(*controller.BookingRequestControllerImpl)))

var bookingRequestRepoSet = wire.NewSet(repository.BookingRequestRepositoryInit,
	wire.Bind(new(repository.BookingRequestRepository), new(*repository.BookingRequestRepositoryImpl)))

func Init() *Initialize {
	wire.Build(
		db,
		authCtrlSet,
		authSvcSet,
		userCtrlSet,
		userSvcSet,
		userRepoSet,
		roomCtrlSet,
		roomSvcSet,
		roomRepoSet,
		geographyCtrlSet,
		geographySvcSet,
		geographyRepoSet,
		addressCtrlSet,
		addressSvcSet,
		addressRepoSet,
		serviceCtrlSet,
		serviceSvcSet,
		serviceRepoSet,
		hashContractSvcSet,
		hashContractRepoSet,
		contractCtrlSet,
		contractSvcSet,
		contractRepoSet,
		invoiceCtrlSet,
		invoiceSvcSet,
		invoiceRepoSet,
		servicesDemandCtrlSet,
		servicesDemandSvcSet,
		servicesDemandRepoSet,
		bookingRequestCtrlSet,
		bookingRequestSvcSet,
		bookingRequestRepoSet,
		NewInitialize)
	return nil
}
