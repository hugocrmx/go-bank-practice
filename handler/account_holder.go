package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/data"
)

func SearchAllAccounts(c *gin.Context) {
	tx := data.DB.MustBegin()
	accounts, err := data.GetAllAccountHolders(tx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error:": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"Accounts": accounts,
	})

}

func SearchRandomccounts(c *gin.Context) {
	tx := data.DB.MustBegin()
	account, err := data.GetRandomAccount(tx)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error:": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"Account": account,
	})

}
