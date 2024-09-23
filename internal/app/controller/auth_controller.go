package controller

import (
	"awesomeProject/internal/app/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
}

type AuthControllerImpl struct {
	authService service.AuthService
}

// Login godoc
// @Summary Login
// @Description Login
// @Tags auth
// @Accept mpfd
// @Produce json
// @Param email formData string true "email"
// @Param password formData string true "password"
// @Success 200 {object} service.LoginResponse
// @Router /auth/login [post]
func (a AuthControllerImpl) Login(c *gin.Context) {
	a.authService.Login(c)
}

func (a AuthControllerImpl) Register(c *gin.Context) {
	a.authService.Register(c)
}

func AuthControllerInit(authService service.AuthService) *AuthControllerImpl {
	return &AuthControllerImpl{authService: authService}
}
