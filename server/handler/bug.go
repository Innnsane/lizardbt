package handler

import (
	"github.com/gin-gonic/gin"
	"lizardbt/global"
	"lizardbt/model/conv"
	"lizardbt/model/conv/s2r"
	"lizardbt/service"
	"net/http"
)

func init() {
	bugGroup := global.Gin.Group("/bug")
	bugGroup.GET("/list", BugList)
}

func BugList(c *gin.Context) {
	var err error
	schemas, err := service.QueryBugs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
	}
	response := conv.List(schemas, s2r.BugBrief)
	c.JSON(http.StatusOK, response)
}
