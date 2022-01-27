package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/handler"
)

func CreateProductSegmentRoutes(g *gin.Engine) {
	g.GET("/api/v1/segments", handler.SearchAllSegments)
	g.GET("/api/v1/segment", handler.SearchRandomSegment)
}
