package service

import (
	"awesomeProject/internal/app/auth"
	"awesomeProject/internal/app/constant"
	"awesomeProject/internal/app/domain/dao"
	"awesomeProject/internal/app/pkg"
	"awesomeProject/internal/app/repository"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type AuthService interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthServiceImpl struct {
	userRepo repository.UserRepository
}

type LoginResponse struct {
	AccessToken  string            `json:"access_token"`
	RefreshToken string            `json:"refresh_token"`
	UserInfo     dao.UsersResponse `json:"user_info"`
}

type LoginAndRegisterParams struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (authSvc AuthServiceImpl) Login(c *gin.Context) {
	var params LoginAndRegisterParams
	if err := c.Bind(&params); err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err.Error(), pkg.Null()))
		return
	}
	data, err := authSvc.userRepo.GetByPhone(params.Phone)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, pkg.BuildResponse(constant.DataNotFound, "Account not found", pkg.Null()))
			return
		}
	}

	if pkg.ComparePassword(params.Password, data.Password) {
		tokenString, err := auth.GenerateJWT(params.Phone, string(data.Role))
		if err != nil {
			c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.Success, "Error creating token", data))
			return
		}
		response := LoginResponse{AccessToken: tokenString, RefreshToken: "", UserInfo: *data}
		c.JSON(http.StatusOK, pkg.BuildResponse(constant.Success, pkg.Null(), response))
	} else {
		c.JSON(http.StatusUnauthorized, pkg.BuildResponse(constant.Unauthorized, "Unauthorized", pkg.Null()))
	}
}

func (authSvc AuthServiceImpl) Register(c *gin.Context) {
	var params LoginAndRegisterParams
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, err.Error(), pkg.Null()))
		return
	}

	_, err := authSvc.userRepo.GetByPhone(params.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			hashedPassword, _ := pkg.HashPassword(params.Password)
			params.Password = hashedPassword
			user := &dao.Users{
				Phone:    params.Phone,
				Password: params.Password,
				Role:     constant.USER_ROLE,
			}
			_, err := authSvc.userRepo.Create(user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, pkg.BuildResponse(constant.InternalServerError, "Error creating user", pkg.Null()))
				return
			}
			c.JSON(http.StatusCreated, pkg.BuildResponse(constant.Success, "User created", user))
		}
	} else {
		c.JSON(http.StatusBadRequest, pkg.BuildResponse(constant.BadRequest, "User already exists", pkg.Null()))
	}
}

func AuthServiceInit(userRepo repository.UserRepository) *AuthServiceImpl {
	return &AuthServiceImpl{userRepo}
}
