package site

import "github.com/gin-gonic/gin"

type siteRoute struct {
	route   *gin.Engine
	handler SiteHandler
}

type SiteRoute interface {
	Init()
}

func NewSiteRoute(route *gin.Engine, handler SiteHandler) SiteRoute {
	return &siteRoute{route, handler}

}

func (s *siteRoute) Init() {
	s.route.POST("/", s.handler.Create())
	s.handler.Create()
}
