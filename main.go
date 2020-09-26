package main

import (
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/v1/account/add", controller.CreateAccount)
	router.POST("/api/v1/login",middleware.Auth,controller.GetAccount)
	router.GET("/api/v1/account/:idAccount", controller.GetAccount)
	router.POST("/api/v1/transfer",controller.Transfer)
	router.POST("/api/v1/withdraw",controller.Withdraw)
	router.POST("/api/v1/deposit",controller.Deposit)
	router.Run(":8080")
}

