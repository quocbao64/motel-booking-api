package service

import (
	"awesomeProject/internal/app/blockchain"
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"time"
)

type ContractService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Liquidity(c *gin.Context)
	CreateFromBookingRequest(c *gin.Context)
}

type ContractServiceImpl struct {
	contractRepo        repository.ContractRepository
	hashContractRepo    repository.HashContractRepository
	servicesDemandRepo  repository.ServicesDemandRepository
	invoiceRepo         repository.InvoiceRepository
	roomRepo            repository.RoomRepository
	userRepo            repository.UserRepository
	servicesHistoryRepo repository.ServicesHistoryRepository
	borrowedItemRepo    repository.BorrowedItemRepository
	transactionRepo     repository.TransactionRepository
	bookingRequestRepo  repository.BookingRequestRepository
}

type ContractParams struct {
	RenterID           int       `json:"renter_id"`
	LessorID           int       `json:"lessor_id"`
	RoomID             int       `json:"room_id"`
	StartDate          time.Time `json:"start_date"`
	DatePay            time.Time `json:"date_pay"`
	PayMode            int       `json:"pay_mode"`
	Payment            float64   `json:"payment"`
	Status             int       `json:"status"`
	IsEnable           bool      `json:"is_enable"`
	FileBase64         string    `json:"file_base64"`
	FileName           string    `json:"file_name"`
	ChargeableServices []int     `json:"chargeable_services"`
	IsRenterSigned     bool      `json:"is_renter_signed"`
	IsLessorSigned     bool      `json:"is_lessor_signed"`
	MonthlyPrice       float64   `json:"monthly_price"`
	BorrowedItems      []int     `json:"borrowed_items"`
	Deposit            float64   `json:"deposit"`
	RentalDuration     int       `json:"rental_duration"`
	Title              string    `json:"title"`
}

type LiquidityParams struct {
	ContractID    int    `json:"contract_id"`
	BorrowedItems []uint `json:"borrowed_items"`
}

type ContractFromBookingRequestParams struct {
	BookingRequestID int    `json:"booking_request_id"`
	PayFor           int    `json:"pay_for"`
	FileBase64       string `json:"file_base64"`
	FileName         string `json:"file_name"`
}

func (repo ContractServiceImpl) GetAll(c *gin.Context) {
	renterID, _ := strconv.Atoi(c.Query("renter_id"))
	lessorID, _ := strconv.Atoi(c.Query("lessor_id"))

	filter := &repository.ContractFilter{
		RenterID: renterID,
		LessorID: lessorID,
	}

	data, err := repo.contractRepo.GetAll(filter)

	//contract := blockchain.InvokeChaincode()
	//contracts, err := blockchain.GetAllContract(contract)
	//
	//fmt.Println(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	var contracts []dao.ContractResponse
	for _, contract := range data {
		renter, _ := repo.userRepo.GetByID(int(contract.RenterID))
		lessor, _ := repo.userRepo.GetByID(int(contract.LessorID))
		room, _ := repo.roomRepo.GetByID(int(contract.RoomID))
		var canceledBy *dao.UsersResponse
		if contract.CanceledBy != nil {
			canceledBy, err = repo.userRepo.GetByID(int(*contract.CanceledBy))
		}
		invoices, err := repo.invoiceRepo.GetByContractID(int(contract.ID))
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}
		contracts = append(contracts, dao.ContractResponse{
			ID:              contract.ID,
			Renter:          *renter,
			Lessor:          *lessor,
			Room:            *room,
			MonthlyPrice:    contract.MonthlyPrice,
			CanceledBy:      canceledBy,
			StartDate:       contract.StartDate,
			DatePay:         contract.DatePay,
			PayMode:         contract.PayMode,
			Payment:         contract.Payment,
			Status:          contract.Status,
			IsEnable:        contract.IsEnable,
			FilePath:        contract.FilePath,
			Invoices:        invoices,
			ServicesHistory: contract.ServicesHistory,
			BorrowedItems:   contract.BorrowedItems,
			Deposit:         contract.Deposit,
			DamagedItems:    contract.DamagedItems,
		})
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), contracts))
}

func (repo ContractServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.contractRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	if data.IsLessorSigned && data.IsRenterSigned {
		contractBC := blockchain.InvokeChaincode()
		_, err = blockchain.ReadAsset(contractBC, int64(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}

		base64, err := pkg.ConvertFileToBase64(data.FilePath)
		if err != nil {
			return
		}

		fileHashed, _ := pkg.HashFileBase64ToSHA256(base64)

		hashContractData, err := repo.hashContractRepo.GetByContractID(int(data.ID))

		if hashContractData == nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), "Hash contract not found"))
			return
		}
		hashContract, err := blockchain.ReadHashContract(contractBC, int64(hashContractData.ID))

		if hashContract != nil && fileHashed != hashContract.Hash {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), "File is modified"))
			return
		}
	}

	renter, _ := repo.userRepo.GetByID(int(data.RenterID))
	lessor, _ := repo.userRepo.GetByID(int(data.LessorID))
	room, _ := repo.roomRepo.GetByID(int(data.RoomID))
	var canceledBy *dao.UsersResponse
	if data.CanceledBy != nil {
		canceledBy, err = repo.userRepo.GetByID(int(*data.CanceledBy))
	}
	invoices, err := repo.invoiceRepo.GetByContractID(int(data.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}
	contract := dao.ContractResponse{
		ID:              data.ID,
		Renter:          *renter,
		Lessor:          *lessor,
		Room:            *room,
		MonthlyPrice:    data.MonthlyPrice,
		CanceledBy:      canceledBy,
		StartDate:       data.StartDate,
		DatePay:         data.DatePay,
		PayMode:         data.PayMode,
		Payment:         data.Payment,
		Status:          data.Status,
		IsEnable:        data.IsEnable,
		FilePath:        data.FilePath,
		Invoices:        invoices,
		ServicesHistory: data.ServicesHistory,
		BorrowedItems:   data.BorrowedItems,
		Deposit:         data.Deposit,
		DamagedItems:    data.DamagedItems,
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), contract))
}

func (repo ContractServiceImpl) Create(c *gin.Context) {
	var contractDAO ContractParams
	err := c.BindJSON(&contractDAO)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	fileContent, err := base64.StdEncoding.DecodeString(contractDAO.FileBase64)
	url, err := pkg.UploadS3("rooms/"+uuid.New().String()+"/"+contractDAO.FileName, fileContent, "application/pdf")
	if err != nil {
		return
	}

	filter := &repository.BorrowedItemFilter{
		IDs: contractDAO.BorrowedItems,
	}
	borrowedItems, err := repo.borrowedItemRepo.GetAll(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	contract := &dao.Contract{
		RenterID:       uint(contractDAO.RenterID),
		LessorID:       uint(contractDAO.LessorID),
		RoomID:         uint(contractDAO.RoomID),
		StartDate:      contractDAO.StartDate,
		DatePay:        contractDAO.DatePay,
		PayMode:        contractDAO.PayMode,
		Payment:        contractDAO.Payment,
		Status:         contractDAO.Status,
		IsEnable:       contractDAO.IsEnable,
		FilePath:       url,
		IsRenterSigned: contractDAO.IsRenterSigned,
		IsLessorSigned: contractDAO.IsLessorSigned,
		MonthlyPrice:   contractDAO.MonthlyPrice,
		BorrowedItems:  borrowedItems,
		Deposit:        contractDAO.Deposit,
		RentalDuration: contractDAO.RentalDuration,
		Title:          contractDAO.Title,
	}
	data, err := repo.contractRepo.Create(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	room, err := repo.roomRepo.GetByID(contractDAO.RoomID)
	var servicesHistory []dao.ServicesHistory
	for _, service := range room.Services {
		servicesHistory = append(servicesHistory, dao.ServicesHistory{
			ServiceID:   service.ID,
			IconURL:     service.IconURL,
			Price:       service.Price,
			IsEnable:    service.IsEnable,
			ContractID:  data.ID,
			ServiceName: service.Name,
		})
	}

	dataHistory, err := repo.servicesHistoryRepo.CreateMultiple(servicesHistory)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	if len(dataHistory) > 0 {
		data.ServicesHistory = dataHistory
	}

	_, err = repo.contractRepo.Update(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ContractServiceImpl) Update(c *gin.Context) {
	var contract *dao.Contract
	err := c.BindJSON(&contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	data, err := repo.contractRepo.Update(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	if data.IsLessorSigned && data.IsRenterSigned {
		file, err := pkg.GetFileFromS3(data.FilePath, "application/pdf")
		hashFile, _ := pkg.HashFileBase64ToSHA256(base64.StdEncoding.EncodeToString(file))
		hashContract := &dao.HashContract{
			ContractID: data.ID,
			Hash:       hashFile,
		}
		dataHash, err := repo.hashContractRepo.Create(hashContract)

		contractBC := blockchain.InvokeChaincode()
		err = blockchain.CreateAssets(contractBC, data)

		if err != nil {
			_ = repo.contractRepo.Delete(int(data.ID))
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}

		err = blockchain.CreateHashContract(contractBC, dataHash)

		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ContractServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.contractRepo.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func (repo ContractServiceImpl) Liquidity(c *gin.Context) {
	params := &LiquidityParams{}
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	contract, err := repo.contractRepo.GetByID(params.ContractID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	var refund float64
	var damagedItems []*dao.DamagedItem
	for _, borrowedItemID := range params.BorrowedItems {
		for _, borrowedItem := range contract.BorrowedItems {
			if borrowedItem.ID == borrowedItemID {
				refund += borrowedItem.Price
				damagedItem := &dao.DamagedItem{
					BaseModel: dao.BaseModel{
						ID: borrowedItemID,
					},
					ContractID:       borrowedItem.ContractID,
					Name:             borrowedItem.Name,
					Price:            borrowedItem.Price,
					RoomID:           borrowedItem.RoomID,
					BookingRequestID: borrowedItem.BookingRequestID,
				}
				damagedItems = append(damagedItems, damagedItem)
			}
		}
	}

	lessorTrans := &dao.Transaction{
		UserID:          contract.LessorID,
		TransactionType: constant.TRANSACTION_REFUND,
		Amount:          contract.Deposit + refund,
		Status:          constant.TRANSACTION_SUCCESS,
		TransactionNo:   uuid.New().String(),
	}

	renterTrans := &dao.Transaction{
		UserID:          contract.RenterID,
		TransactionType: constant.TRANSACTION_REFUND,
		Amount:          contract.Deposit - refund,
		Status:          constant.TRANSACTION_SUCCESS,
		TransactionNo:   uuid.New().String(),
	}

	lessorTransCreated, err := repo.transactionRepo.Create(lessorTrans)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	renterTransCreated, err := repo.transactionRepo.Create(renterTrans)
	if err != nil {
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	err = repo.borrowedItemRepo.CreateDamagedItems(damagedItems)

	if err != nil {
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		_ = repo.transactionRepo.Delete(int(renterTransCreated.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	resp, err := repo.contractRepo.UpdateLiquidity(contract.ID, lessorTrans, renterTrans)

	if err != nil {
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		_ = repo.transactionRepo.Delete(int(renterTransCreated.ID))
		_ = repo.borrowedItemRepo.Delete(int(contract.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), resp))
}

func (repo ContractServiceImpl) CreateFromBookingRequest(c *gin.Context) {
	params := &ContractFromBookingRequestParams{}
	err := c.BindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	bookingRequest, err := repo.bookingRequestRepo.GetByID(params.BookingRequestID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	room, err := repo.roomRepo.GetByID(int(bookingRequest.RoomID))

	var payment float64
	if params.PayFor == 1 {
		payment = room.Deposit
	} else {
		payment = room.Deposit + room.Price
	}

	renter, err := repo.userRepo.GetByID(int(bookingRequest.RenterID))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}
	lessor, err := repo.userRepo.GetByID(int(bookingRequest.LessorID))
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	if lessor.Balance < payment {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, "Lessor balance is not enough", pkg.Null()))
		return
	}

	if renter.Balance < payment {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, "Renter balance is not enough", pkg.Null()))
		return
	}

	fileContent, err := base64.StdEncoding.DecodeString(params.FileBase64)
	url, err := pkg.UploadS3("rooms/"+uuid.New().String()+"/"+params.FileName, fileContent, "application/pdf")
	if err != nil {
		return
	}

	contract := &dao.Contract{
		RenterID:       bookingRequest.RenterID,
		LessorID:       bookingRequest.LessorID,
		RoomID:         bookingRequest.RoomID,
		StartDate:      bookingRequest.StartDate,
		DatePay:        time.Now(),
		PayMode:        constant.PAYMODE_VNPAY,
		Payment:        payment,
		Status:         constant.CONTRACT_ACTIVE,
		FilePath:       url,
		IsRenterSigned: true,
		IsLessorSigned: false,
		Deposit:        room.Deposit,
		IsEnable:       true,
	}

	data, err := repo.contractRepo.Create(contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	bookingRequest.Status = constant.BOOKING_REQUEST_ACCEPTED
	bookingRes, err := repo.bookingRequestRepo.Update(bookingRequest)
	if err != nil {
		_ = repo.contractRepo.Delete(int(data.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	lessorTrans := &dao.Transaction{
		UserID:          contract.LessorID,
		TransactionType: constant.TRANSACTION_PAYMENT,
		Amount:          room.Deposit,
		Status:          constant.TRANSACTION_SUCCESS,
		TransactionNo:   uuid.New().String(),
	}

	renterTrans := &dao.Transaction{
		UserID:          contract.RenterID,
		TransactionType: constant.TRANSACTION_PAYMENT,
		Amount:          payment,
		Status:          constant.TRANSACTION_SUCCESS,
		TransactionNo:   uuid.New().String(),
	}

	lessorTransCreated, err := repo.transactionRepo.Create(lessorTrans)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_REQUEST_PROCESSING
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	renterTransCreated, err := repo.transactionRepo.Create(renterTrans)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_REQUEST_PROCESSING
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	invoice := &dao.Invoice{
		ContractID:    &data.ID,
		Amount:        payment,
		PaymentStatus: constant.PAYMENT_COMPLETED,
		TransactionID: renterTransCreated.ID,
	}
	_, err = repo.invoiceRepo.Create(invoice)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_REQUEST_PROCESSING
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		_ = repo.transactionRepo.Delete(int(renterTransCreated.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	err = repo.userRepo.UpdateBalance(lessor.ID, constant.TRANSACTION_PAYMENT, room.Deposit)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_REQUEST_PROCESSING
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		_ = repo.transactionRepo.Delete(int(renterTransCreated.ID))
		_ = repo.invoiceRepo.Delete(int(invoice.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	err = repo.userRepo.UpdateBalance(renter.ID, constant.TRANSACTION_PAYMENT, payment)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_REQUEST_PROCESSING
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		_ = repo.transactionRepo.Delete(int(renterTransCreated.ID))
		_ = repo.invoiceRepo.Delete(int(invoice.ID))
		_ = repo.userRepo.UpdateBalance(lessor.ID, constant.TRANSACTION_REFUND, room.Deposit)
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func ContractServiceInit(
	repo repository.ContractRepository,
	hashContractRepo repository.HashContractRepository,
	servicesDemandRepo repository.ServicesDemandRepository,
	invoiceRepo repository.InvoiceRepository,
	roomRepo repository.RoomRepository,
	userRepo repository.UserRepository,
	servicesHistoryRepo repository.ServicesHistoryRepository,
	borrowedItemRepo repository.BorrowedItemRepository,
	transactionRepo repository.TransactionRepository,
	bookingRequestRepo repository.BookingRequestRepository) *ContractServiceImpl {
	return &ContractServiceImpl{
		contractRepo:        repo,
		hashContractRepo:    hashContractRepo,
		servicesDemandRepo:  servicesDemandRepo,
		invoiceRepo:         invoiceRepo,
		roomRepo:            roomRepo,
		userRepo:            userRepo,
		servicesHistoryRepo: servicesHistoryRepo,
		borrowedItemRepo:    borrowedItemRepo,
		transactionRepo:     transactionRepo,
		bookingRequestRepo:  bookingRequestRepo,
	}
}
