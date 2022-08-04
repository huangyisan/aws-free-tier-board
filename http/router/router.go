package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-dashboard/http/handler/awscost"
	"trojan-dashboard/http/router/middleware"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// MiddleWares
	g.Use(gin.Recovery())
	g.Use(middleware.Options)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	u := g.Group("/v1/cost")
	{
		u.GET("", awscost.List)
	}
	return g
}
