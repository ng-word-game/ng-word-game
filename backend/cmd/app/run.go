package app

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kikils/go-websocket/project/config"
	"github.com/volatiletech/sqlboiler/boil"
)

var router *gin.Engine

func Run() {
	config.InitConfig()

	router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)
	if config.Config.App.Env == "local" {
		router.Use(gin.Logger())
		gin.SetMode(gin.DebugMode)
		boil.DebugMode = false
	}
	router.Use(middleware.Cors())
	router.Use(middleware.Trace())
	router.Use(middleware.AppVersion())

	initRouters(authH, userH, videoH, relH, middleware.Authenticate())
	err := router.Run(":" + config.Config.App.Port)
	if err != nil {
		log.Println(err)
	}
}

func initRouters() {
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "app",
		})
	})
	{
		router.POST("/signin", authMiddleware, ah.Signin)
	}
}
