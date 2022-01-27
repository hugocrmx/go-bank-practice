package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/data"
)

func SearchAllCreditCards(c *gin.Context) {
	tx := data.DB.MustBegin()
	cards, err := data.GetAllCreditCards(tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"cards": cards,
	})
}

func CreateCreditCard(c *gin.Context) {
	var id int64
	var err error

	card := &data.CreditCard{}
	if err = c.ShouldBindJSON(card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := data.DB.MustBegin()
	if id, err = data.AddCreditCard(tx, card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func SearchCreditCardById(c *gin.Context) {
	tx := data.DB.MustBegin()
	creditCaridIdInt, _ := strconv.Atoi(c.Param("id"))
	creditCard, err := data.FindCreditCardById(tx, creditCaridIdInt)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{
		"CreditCard": creditCard,
	})
}
