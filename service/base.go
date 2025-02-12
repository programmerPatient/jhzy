package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseService struct {
}

func (s *BaseService) Success(c *gin.Context, date interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"date": date,
	})
}
func (s *BaseService) HTML(c *gin.Context, name string, data gin.H) {
	c.HTML(http.StatusOK, name, data)
}
