package controller

import (
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/model"
	"github.com/FadhlanHawali/Digitalent-Kominfo_Implementation-MVC-Golang/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (inDB *InDB) CreateAccount (c *gin.Context){

	var account model.Account
	if err := c.Bind(&account); err!= nil {
		utils.WrapAPIError(c,err.Error(),http.StatusBadRequest)
		return
	}



}

func (inDB *InDB) GetAccount (c *gin.Context){

}

func (inDB *InDB) Deposit (c *gin.Context){

}

func (inDB *InDB) Withdraw (c *gin.Context){

}
