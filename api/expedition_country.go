package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/handler"
)

func CreateExpeditionCountryRoutes(e *gin.Engine) {
	e.GET("/api/v1/country", handler.SearchRandomCountry)
}
