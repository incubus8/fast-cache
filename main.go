package main

import (
	"github.com/gin-gonic/gin"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/hashicorp/golang-lru"
	"io/ioutil"
	"github.com/pkg/errors"
)

type Application struct {
	cache *lru.Cache
}

var app *Application

func NewApplication() {
	app = &Application{}
	app.cache, _ = lru.New(10000)

	g.Run(g.Config{
		ListenAddr: "localhost:8080",
		Handler: app.router(),
	})
}

func main() {
	NewApplication()
}

func (app *Application) router() *gin.Engine {
	router := gin.New()

	api := router.Group("api")
	{
		v1 := api.Group("v1")
		{
			v1.GET("/lru/:key", app.GetCache)
			v1.DELETE("/lru/:key", app.DelCache)
			v1.POST("/lru/:key", app.AddCache)
			v1.PUT("/lru/:key", app.AddCache)
		}
	}
	return router
}

func (app *Application) AddCache(c *gin.Context) {
	if data, err := ioutil.ReadAll(c.Request.Body); err != nil {
		c.AbortWithError(500, err)
		return
	} else {
		key := c.Param("key")
		if key == "" {
			c.AbortWithError(500, errors.New("invalid key"))
			return
		}

		app.cache.Add(key, data)
		c.JSON(200, string(data))
	}
}

func (app *Application) GetCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
	}

	data, ok := app.cache.Get(key)
	if !ok {
		c.AbortWithError(500, errors.New("no lru cache with this key: " + key))
		return
	}

	c.JSON(200, string(data.([]byte)))
}

func (app *Application) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
	}

	app.cache.Remove(key)

	c.Status(204)
}