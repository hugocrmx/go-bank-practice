package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/handler"
)

func CreateCreditCardRoutes(g *gin.Engine) {
	g.GET("/api/v1/cards", handler.SearchAllCreditCards)
	g.POST("/api/v1/cards", handler.CreateCreditCard)
	g.GET("/api/v1/cards/:id", handler.SearchCreditCardById)
}
