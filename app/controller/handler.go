package controller

import (
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/model"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateAccount (c *gin.Context){

	var account model.Account
	if err := c.Bind(&account); err!= nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
	flag,err := model.InsertNewAccount(account)
	if flag{
		utils.WrapAPISuccess(c,"success",http.StatusOK)
		return
	}else {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
}

func GetAccount (c *gin.Context){
	idAccount := c.Param("idAccount")
	idAccountInt,err := strconv.Atoi(idAccount); if err != nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}

	flag,err,resp := model.GetAccountDetail(idAccountInt);if err != nil {
		utils.WrapAPIError(c,err.Error(),http.StatusInternalServerError)
		return
	}

	if flag{
		utils.WrapAPIData(c,resp,http.StatusOK,"success")
		return
	}
}

func Transfer (c *gin.Context){
	var transaction model.Transaction
	if err := c.Bind(&transaction); err!= nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}

	flag,err := model.Transfer(transaction); if flag{
		utils.WrapAPISuccess(c,"success",http.StatusOK)
		return
	}else {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
}

func Withdraw (c *gin.Context){
	var transaction model.Transaction
	if err := c.Bind(&transaction); err!= nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}

	flag,err := model.Withdraw(transaction); if flag{
		utils.WrapAPISuccess(c,"success",http.StatusOK)
		return
	}else {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
}

func Deposit (c *gin.Context){
	var transaction model.Transaction
	if err := c.Bind(&transaction); err!= nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}

	flag,err := model.Deposit(transaction); if flag{
		utils.WrapAPISuccess(c,"success",http.StatusOK)
		return
	}else {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
}