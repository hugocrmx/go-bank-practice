package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/handler"
)

func CreateBankRoutes(g *gin.Engine) {
	g.GET("/api/v1/banks", handler.SearchAllBanks)
	g.GET("/api/v1/banks/random", handler.SearchRandomBank)
}
