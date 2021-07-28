package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jasonnchann24/gogin-rest/controllers"
	"github.com/jasonnchann24/gogin-rest/middlewares"
	"github.com/jasonnchann24/gogin-rest/services"
	// gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService services.VideoService = services.New()
	loginService services.LoginService = services.NewLoginService()
	jwtService   services.JWTService   = services.NewJWTService()

	videoController controllers.VideoController = controllers.New(videoService)
	loginController controllers.LoginController = controllers.NewLoginController(loginService, jwtService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	// server := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery(), middlewares.Logger())
	// middlewares.BasicAuth(), gindump.Dump()

	server.Static("/css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{
					"message": "Video saved",
				})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
