package main

import (
	"fmt"
	"io"
	"os"

	"./common"
	"./db"
	apis "./rest"
	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "./docs"
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

	// Initialize User database
	err = db.Database.Init()
	if err != nil {
		fmt.Println(err)
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

// @title User Service API Document
// @version 1.0
// @description List APIs of User Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8080
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
		admin := v1.Group("/admin")
		{
			admin.POST("/auth", c.Authenticate)
		}

		user := v1.Group("/user")

		// APIs need to be authenticated
		user.Use(jwt.Auth(common.Config.JwtSecretPassword))
		{
			user.POST("", c.AddUser)
			user.GET("/list", c.ListUsers)
			user.GET("detail/:id", c.GetUserByID)
			user.GET("/", c.GetUserByParams)
			user.DELETE(":id", c.DeleteUserByID)
			user.PATCH("", c.UpdateUser)
		}
	}

	m.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	m.router.Run(common.Config.Port)
}
