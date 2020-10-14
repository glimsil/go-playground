package main

import (
	"io"
	"os"

	"./apis"
	"./common"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Main struct {
	router *gin.Engine
}

func (m *Main) initServer() error {
	var err error
	// Load config file
	err = common.LoadConfig()
	if err != nil {
		return err
	}

	// Setting Gin Logger
	if common.Config.EnableGinFileLog {
		f, _ := os.Create("logs/gin.log")
		if common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter(os.Stdout, f)
		} else {
			gin.DefaultWriter = io.MultiWriter(f)
		}
	} else {
		if !common.Config.EnableGinConsoleLog {
			gin.DefaultWriter = io.MultiWriter()
		}
	}

	m.router = gin.Default()

	return nil
}

// @title Test Service API Document
// @version 1.0
// @description List APIs of Test Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8808
// @BasePath /api/v1
func main() {
	m := Main{}

	// Initialize server
	if m.initServer() != nil {
		return
	}

	c := apis.User{}
	// Simple group: v1
	v1 := m.router.Group("/api/v1")
	{
		user := v1.Group("/user")
		user.GET("/", c.GetUser)
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
