package rest

import (
	"github.com/gin-gonic/gin"
)

func (app *Application) router() *gin.Engine {
	router := gin.New()

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/lru/:key", app.lruApp.GetCache)
			v1.DELETE("/lru/:key", app.lruApp.DelCache)
			v1.POST("/lru/:key", app.lruApp.AddCache)
			v1.PUT("/lru/:key", app.lruApp.AddCache)

			v1.GET("/arc/:key", app.arcApp.GetCache)
			v1.DELETE("/arc/:key", app.arcApp.DelCache)
			v1.POST("/arc/:key", app.arcApp.AddCache)
			v1.PUT("/arc/:key", app.arcApp.AddCache)

			v1.GET("/ccache/:key", app.ccacheApp.GetCache)
			v1.DELETE("/ccache/:key", app.ccacheApp.DelCache)
			v1.POST("/ccache/:key", app.ccacheApp.AddCache)
			v1.PUT("/ccache/:key", app.ccacheApp.AddCache)
		}
	}

	router.GET("/stats", statsHandler)
	router.Any("/metrics", gin.WrapH(app.GetPrometheusExporter()))

	return router
}
