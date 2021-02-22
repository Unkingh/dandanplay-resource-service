package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"dandanplay-resource-service/api"
	"dandanplay-resource-service/api/dmhy"
	"dandanplay-resource-service/config"
)

var r *gin.Engine

// InitRouter creates a gin.Engine instance
func InitRouter() *gin.Engine {
	r = gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, config.HtmlStringIndex)
	})

	register(dmhy.Provider)

	return r
}

// register adds a new router group according to the settings of api.Provider
func register(p *api.Provider) {
	if !p.IsEnabled {
		return
	}

	group := r.Group(p.Route)
	{
		group.GET("/type", p.GenerateType)
		group.GET("/subgroup", p.GenerateSubgroup)
		group.GET("/list", p.GenerateList)
	}
}
