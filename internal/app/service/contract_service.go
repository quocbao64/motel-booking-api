package service

import (
	"awesomeProject/internal/app/blockchain"
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
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
	CancelContract(c *gin.Context)
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

type CancelContractParams struct {
	ContractID   int  `json:"contract_id"`
	CancelStatus int  `json:"cancel_status"`
	CanceledBy   uint `json:"canceled_by"`
	Status       int  `json:"status"`
}

func (repo ContractServiceImpl) GetAll(c *gin.Context) {
	renterID, _ := strconv.Atoi(c.Query("renter_id"))
	lessorID, _ := strconv.Atoi(c.Query("lessor_id"))

	// Initialize the filter to pass to the blockchain query
	filter := &repository.ContractFilter{
		RenterID: renterID,
		LessorID: lessorID,
	}

	// Get the contract data from the blockchain
	contract := blockchain.InvokeChaincode()
	contractBC, err := blockchain.ReadAllContracts(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	// Create a slice to hold the filtered contracts
	var contracts []dao.ContractResponse

	// Filter contracts based on renterID and lessorID
	for _, bcContract := range contractBC {
		// Apply filter logic
		if (filter.RenterID != 0 && bcContract.RenterID != uint(filter.RenterID)) ||
			(filter.LessorID != 0 && bcContract.LessorID != uint(filter.LessorID)) {
			// Skip contracts that don't match the filter criteria
			continue
		}

		// Use the blockchain data to populate the contract response
		renter, _ := repo.userRepo.GetByID(int(bcContract.RenterID))
		lessor, _ := repo.userRepo.GetByID(int(bcContract.LessorID))
		room, _ := repo.roomRepo.GetByID(int(bcContract.RoomID))

		var canceledBy *dao.UsersResponse
		if bcContract.CanceledBy != nil {
			canceledBy, err = repo.userRepo.GetByID(int(*bcContract.CanceledBy))
		}

		// Retrieve invoices from the database based on the contract ID
		invoices, err := repo.invoiceRepo.GetByContractID(int(bcContract.ID))
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}

		// Append the contract response data from blockchain to the slice
		contracts = append(contracts, dao.ContractResponse{
			ID:              bcContract.ID,
			Renter:          *renter,
			Lessor:          *lessor,
			Room:            *room,
			MonthlyPrice:    bcContract.MonthlyPrice,
			CanceledBy:      canceledBy,
			StartDate:       bcContract.StartDate,
			DatePay:         bcContract.DatePay,
			PayMode:         bcContract.PayMode,
			Payment:         bcContract.Payment,
			Status:          bcContract.Status,
			IsEnable:        bcContract.IsEnable,
			FilePath:        bcContract.FilePath,
			Invoices:        invoices,
			ServicesHistory: bcContract.ServicesHistory,
			BorrowedItems:   bcContract.BorrowedItems,
			Deposit:         bcContract.Deposit,
			DamagedItems:    bcContract.DamagedItems,
			RentalDuration:  bcContract.RentalDuration,
			CancelStatus:    bcContract.CancelStatus,
		})
	}

	// Return the response with filtered contract data from blockchain
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
		_, err = blockchain.ReadContract(contractBC, int64(id))
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
		RentalDuration:  data.RentalDuration,
		CancelStatus:    data.CancelStatus,
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

	// Decode the base64 file and upload it to S3
	fileContent, err := base64.StdEncoding.DecodeString(contractDAO.FileBase64)
	url, err := pkg.UploadS3("rooms/"+uuid.New().String()+"/"+contractDAO.FileName, fileContent, "application/pdf")
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.InternalServerError, pkg.Null(), err))
		return
	}

	// Get borrowed items
	filter := &repository.BorrowedItemFilter{IDs: contractDAO.BorrowedItems}
	borrowedItems, err := repo.borrowedItemRepo.GetAll(filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	// Create the contract object
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

	// Save contract to the blockchain
	data, err := repo.contractRepo.Create(contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	// Generate a hash for the contract data
	hashData := fmt.Sprintf("%s-%s", contract.ID, contract.FilePath)
	hash, _ := GenerateHash(hashData)

	// Create a `HashContract` object
	hashContract := &dao.HashContract{
		ContractID: data.ID,
		Hash:       hash,
	}

	// Save the HashContract to the blockchain
	_, err = repo.hashContractRepo.Create(hashContract)
	if err != nil {
		c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.InternalServerError, pkg.Null(), err))
		return
	}

	// Create services history
	room, err := repo.roomRepo.GetByID(contractDAO.RoomID)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}
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

	// Update contract with services history
	_, err = repo.contractRepo.Update(data)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	// Return success response with contract data
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ContractServiceImpl) Update(c *gin.Context) {
	var contract *dao.Contract
	// Bind the incoming request data to the contract object
	err := c.BindJSON(&contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	// Update the contract in the repository
	data, err := repo.contractRepo.Update(contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	// If both lessor and renter have signed, proceed with hashing and blockchain operations
	if data.IsLessorSigned && data.IsRenterSigned {
		// Retrieve the contract file from S3
		file, err := pkg.GetFileFromS3(data.FilePath, "application/pdf")
		if err != nil {
			c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.InternalServerError, pkg.Null(), err))
			return
		}

		// Generate the hash from the file content
		hashFile, err := pkg.HashFileBase64ToSHA256(base64.StdEncoding.EncodeToString(file))
		if err != nil {
			c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.InternalServerError, pkg.Null(), err))
			return
		}

		// Create a new hash contract record
		hashContract := &dao.HashContract{
			ContractID: data.ID,
			Hash:       hashFile,
		}
		dataHash, err := repo.hashContractRepo.Create(hashContract)
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}

		// Interact with blockchain for contract asset creation
		contractBC := blockchain.InvokeChaincode()
		err = blockchain.CreateContract(contractBC, data)
		if err != nil {
			// If blockchain interaction fails, delete the contract and return an error
			_ = repo.contractRepo.Delete(int(data.ID))
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}

		// Create the hash contract on the blockchain
		err = blockchain.CreateHashContract(contractBC, dataHash)
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}
	}

	// Respond with the updated contract data
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
		MonthlyPrice:   room.Price,
		BorrowedItems:  bookingRequest.BorrowedItems,
		Title:          room.Title,
		RentalDuration: bookingRequest.RentalDuration,
	}

	fmt.Println("Borrowed Items: ", bookingRequest.BorrowedItems)

	data, err := repo.contractRepo.Create(contract)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

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

	if len(servicesHistory) != 0 {
		dataHistory, err := repo.servicesHistoryRepo.CreateMultiple(servicesHistory)
		if err != nil {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
			return
		}

		if len(dataHistory) > 0 {
			data.ServicesHistory = dataHistory
		}

		_, err = repo.contractRepo.Update(data)
	}
	bookingRequest.Status = constant.BOOKING_FINISHED
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
		bookingRequest.Status = constant.BOOKING_WAITING_PAYMENT
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	renterTransCreated, err := repo.transactionRepo.Create(renterTrans)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_WAITING_PAYMENT
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
		bookingRequest.Status = constant.BOOKING_WAITING_PAYMENT
		_, _ = repo.bookingRequestRepo.Update(bookingRes)
		_ = repo.contractRepo.Delete(int(data.ID))
		_ = repo.transactionRepo.Delete(int(lessorTransCreated.ID))
		_ = repo.transactionRepo.Delete(int(renterTransCreated.ID))
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err, pkg.Null()))
		return
	}

	err = repo.userRepo.UpdateBalance(lessor.ID, constant.TRANSACTION_PAYMENT, room.Deposit)
	if err != nil {
		bookingRequest.Status = constant.BOOKING_WAITING_PAYMENT
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
		bookingRequest.Status = constant.BOOKING_WAITING_PAYMENT
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

func (repo ContractServiceImpl) CancelContract(c *gin.Context) {
	params := &CancelContractParams{}
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

	if params.Status == 5 || params.Status == 6 {
		err := repo.roomRepo.UpdateStatus(int(contract.RoomID), constant.ROOM_AVAILABLE)
		if err != nil {
			return
		}
	}

	contract.CancelStatus = params.CancelStatus
	contract.Status = params.Status
	contract.CanceledBy = &params.CanceledBy

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), contract))
}

func GenerateHash(fileBase64 string) (string, error) {
	// Decode the base64 string to get the raw file content
	fileContent, err := base64.StdEncoding.DecodeString(fileBase64)
	if err != nil {
		return "", errors.New("failed to decode base64 string")
	}

	// Generate the SHA-256 hash of the file content
	hash := sha256.New()
	hash.Write(fileContent)
	hashBytes := hash.Sum(nil)

	// Convert the hash bytes to a hexadecimal string
	hashString := hex.EncodeToString(hashBytes)
	return hashString, nil
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
