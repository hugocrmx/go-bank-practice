package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/data"
)

func SearchRandomCountry(c *gin.Context) {
	tx := data.DB.MustBegin()
	country, err := data.GetRandomExpeditionCountry(tx)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Country": country,
	})
}
