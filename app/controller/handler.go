package controller

import (
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/model"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateAccount (c *gin.Context){

	var account model.Account
	if err := c.Bind(&account); err!= nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
	pass, err := utils.HashGenerator(account.Password); if err != nil{
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
	account.Password = pass
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
	idAccount := c.MustGet("account_number").(int)
	flag,err,trx,acc := model.GetAccountDetail(idAccount);if err != nil {
		utils.WrapAPIError(c,err.Error(),http.StatusInternalServerError)
		return
	}
	if flag{
		utils.WrapAPIData(c,map[string]interface{}{
			"account":acc,
			"transaction":trx,
		},http.StatusOK,"success")
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

func Login(c *gin.Context){
	var auth model.Auth
	if err := c.Bind(&auth); err != nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}
	log.Println("LOGIN")
	flag,err,token := model.Login(auth); if flag{
		utils.WrapAPIData(c,map[string]interface{}{
			"token":token,
		},http.StatusOK,"success")
	}else {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
	}
}