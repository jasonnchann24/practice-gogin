package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jasonnchann24/gogin-rest/dto"
	"github.com/jasonnchann24/gogin-rest/services"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jWtService   services.JWTService
}

func NewLoginController(loginService services.LoginService,
	jWtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials dto.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return ""
	}
	isAuthenticated := controller.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return controller.jWtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
