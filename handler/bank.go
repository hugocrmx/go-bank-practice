package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/data"
)

func SearchAllBanks(c *gin.Context) {
	tx := data.DB.MustBegin()
	banks, err := data.GetAllBanks(tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error handler": err.Error(),
		})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"banks": banks,
	})
}

func SearchRandomBank(c *gin.Context) {
	tx := data.DB.MustBegin()
	bank, err := data.GetRandomBank(tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"Bank": bank,
	})
}
