package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/handler"
)

func CreateAccountHolderRoutes(g *gin.Engine) {
	g.GET("/api/v1/accounts", handler.SearchAllAccounts)
	g.GET("/api/v1/account", handler.SearchRandomccounts)
}
