package router

import (
	"url_location/handler/get_url_location"

	"github.com/gin-gonic/gin"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(mw...)

	g.NoRoute(ApiNotFound)
	g.GET("/", ApiHello)
	g.GET("/ping", ApiPing)

	// 路由
	u := g.Group("/api/v1")
	{
		u.GET("/urlOriginAndLocation/", get_url_location.HandlerGetUrlOriginAndLocation)
		u.POST("/urlOriginAndLocation/", get_url_location.HandlerGetUrlOriginAndLocation)

		u.GET("/urlOrigin/", get_url_location.HandlerGetUrlOrigin)
		u.POST("/urlOrigin/", get_url_location.HandlerGetUrlOrigin)

		u.GET("/urlLocation/", get_url_location.HandlerGetUrlLocation)
		u.POST("/urlLocation/", get_url_location.HandlerGetUrlLocation)

	}

	return g
}
