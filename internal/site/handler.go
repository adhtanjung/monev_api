package site

import (
	"github.com/adhtanjung/monev/internal/model"
	"github.com/gin-gonic/gin"
)

type SiteHandler interface {
	Create() gin.HandlerFunc
}

type siteHandler struct {
	repository Repository
}

func NewSiteHandler(repository Repository) SiteHandler {
	return &siteHandler{repository: repository}
}

func (s *siteHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var site model.Site
		c.BindJSON(&site)
		err := s.repository.Create(site)
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
			})
		} else {
			c.JSON(200, gin.H{
				"message": "success",
			})
		}
	}
}
