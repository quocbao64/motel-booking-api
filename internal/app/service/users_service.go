package service

import (
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserService interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	GetByEmail(c *gin.Context)
	GetByPhone(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UserServiceImpl struct {
	userRepo    repository.UserRepository
	addressRepo repository.AddressRepository
}

func (repo UserServiceImpl) GetAll(c *gin.Context) {
	data, err := repo.userRepo.GetAll()

	if err != nil {
		return
	}

	var customerResponse []*dao.UsersResponse
	for _, customer := range data {
		customerResponse = append(customerResponse, customer)
	}
	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), customerResponse))
}

func (repo UserServiceImpl) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, err := repo.userRepo.GetByID(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo UserServiceImpl) GetByEmail(c *gin.Context) {
	email := c.Param("email")
	data, err := repo.userRepo.GetByEmail(email)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func (repo UserServiceImpl) Create(c *gin.Context) {
	var repotomer *dao.Users
	_ = c.BindJSON(&repotomer)

	passwordHashed, _ := pkg.HashPassword(repotomer.Password)

	repotomer.Password = passwordHashed
	repotomerResponse, err := repo.userRepo.Create(repotomer)

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.InvalidRequest, "Email already exist", pkg.Null()))
			return
		}
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), repotomerResponse))
}

func (repo UserServiceImpl) Update(c *gin.Context) {
	var repotomer *dao.Users
	_ = c.BindJSON(&repotomer)

	repotomerResponse, err := repo.userRepo.Update(repotomer)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), repotomerResponse))
}

func (repo UserServiceImpl) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := repo.userRepo.Delete(id)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), pkg.Null()))
}

func (repo UserServiceImpl) GetByPhone(c *gin.Context) {
	phone := c.Param("phone")

	data, err := repo.userRepo.GetByPhone(phone)

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), data))
}

func UserServiceInit(userRepo repository.UserRepository, addressRepo repository.AddressRepository) *UserServiceImpl {
	return &UserServiceImpl{
		userRepo:    userRepo,
		addressRepo: addressRepo,
	}
}
