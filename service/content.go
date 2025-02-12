package service

import (
	"github.com/gin-gonic/gin"
)

type ContentService struct {
	BaseService
}

var ContentS = &ContentService{}

func (s *ContentService) Index(c *gin.Context) {
	s.HTML(c, "content.tmpl", nil)
}

func (s *ContentService) Up(c *gin.Context) {
	s.Success(c, nil)
}
