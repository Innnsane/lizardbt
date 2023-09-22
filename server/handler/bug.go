package handler

import (
	"github.com/gin-gonic/gin"
	"lizardbt/server/global"
	"lizardbt/server/model/conv"
	"lizardbt/server/model/conv/s2r"
	"lizardbt/server/service"
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
