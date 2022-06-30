package main

import (
	"github.com/adhtanjung/monev/database"
	"github.com/adhtanjung/monev/internal/site"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	siteRepo := site.NewSiteRepository(&database.DB)

	siteHandler := site.NewSiteHandler(siteRepo)

	siteRoute := site.NewSiteRoute(r, siteHandler)

	siteRoute.Init()

	r.Run(":8080")
}
