package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugocrmx/go-bank-practice/data"
)

func SearchAllSegments(c *gin.Context) {
	tx := data.DB.MustBegin()
	segments, err := data.GetAllSegments(tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error:": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"segments": segments,
	})
}

func SearchRandomSegment(c *gin.Context) {
	tx := data.DB.MustBegin()
	segment, err := data.GetRandomSegment(tx)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, gin.H{
		"Segment": segment,
	})
}
