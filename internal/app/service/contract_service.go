package service

import (
	"awesomeProject/internal/app/blockchain"
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"fmt"
	"github.com/gin-gonic/gin"
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
	GetAllByRenterOrLessorID(c *gin.Context)
}

type ContractServiceImpl struct {
	contractRepo       repository.ContractRepository
	hashContractRepo   repository.HashContractRepository
	servicesDemandRepo repository.ServicesDemandRepository
	invoiceRepo        repository.InvoiceRepository
}

type ContractParams struct {
	RenterID           int       `json:"renter_id"`
	LessorID           int       `json:"lessor_id"`
	RoomID             int       `json:"room_id"`
	DateRent           time.Time `json:"date_rent"`
	DatePay            time.Time `json:"date_pay"`
	PayMode            int       `json:"pay_mode"`
	Payment            float64   `json:"payment"`
	Status             int       `json:"status"`
	IsEnable           bool      `json:"is_enable"`
	FileBase64         string    `json:"file_base64"`
	ChargeableServices []int     `json:"chargeable_services"`
}

func (repo ContractServiceImpl) GetAll(c *gin.Context) {
	data, err := repo.contractRepo.GetAll()
	fmt.Println("test")

	contract := blockchain.InvokeChaincode()
	blockchain.GetAllAssets(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ContractServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.contractRepo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

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

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo ContractServiceImpl) Create(c *gin.Context) {
	var contractDAO ContractParams
	err := c.BindJSON(&contractDAO)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	filePath, err := pkg.SaveFile(contractDAO.FileBase64)

	contract := &dao.Contract{
		RenterID: uint(contractDAO.RenterID),
		LessorID: uint(contractDAO.LessorID),
		RoomID:   uint(contractDAO.RoomID),
		DateRent: contractDAO.DateRent,
		DatePay:  contractDAO.DatePay,
		PayMode:  contractDAO.PayMode,
		Payment:  contractDAO.Payment,
		Status:   contractDAO.Status,
		IsEnable: contractDAO.IsEnable,
		FilePath: filePath,
	}
	data, err := repo.contractRepo.Create(contract)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	hashFile, _ := pkg.HashFileBase64ToSHA256(contractDAO.FileBase64)
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

func (repo ContractServiceImpl) GetAllByRenterOrLessorID(c *gin.Context) {
	renterID, _ := strconv.Atoi(c.Query("renter_id"))
	lessorID, _ := strconv.Atoi(c.Query("lessor_id"))
	data, err := repo.contractRepo.GetAllByRenterOrLessorID(renterID, lessorID)

	if err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, pkg.Null(), err))
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func ContractServiceInit(
	repo repository.ContractRepository,
	hashContractRepo repository.HashContractRepository,
	servicesDemandRepo repository.ServicesDemandRepository,
	invoiceRepo repository.InvoiceRepository) *ContractServiceImpl {
	return &ContractServiceImpl{
		contractRepo:       repo,
		hashContractRepo:   hashContractRepo,
		servicesDemandRepo: servicesDemandRepo,
		invoiceRepo:        invoiceRepo,
	}
}
