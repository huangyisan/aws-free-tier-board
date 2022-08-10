package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"trojan-dashboard/http/handler/trojan"
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

	traffic := g.Group("/v1/traffic")
	{

		traffic.GET("", trojan.GetAllTraffic)
		traffic.GET("tag/:tag", trojan.GetTrafficByTag)
		traffic.GET("group/:group", trojan.GetTrafficByGroup)
		traffic.GET("group", trojan.GetAllTrafficByGroup)

	}

	servers := g.Group("/v1/servers")
	{
		servers.GET("", trojan.ListServer)
		servers.GET("tag", trojan.ListServerTag)
		servers.GET("group", trojan.ListServerGroup)
		servers.GET("group/:group", trojan.ListServerByGroup)
	}

	return g
}
