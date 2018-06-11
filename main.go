package main

import (
	"github.com/gin-gonic/gin"
	g "github.com/incubus8/go/pkg/gin"
	"github.com/hashicorp/golang-lru"
	"io/ioutil"
	"github.com/pkg/errors"
)

func main() {
	NewApplication()
}

type Application struct {
	lruApp LRUApplication
	arcApp ARCApplication
}

type LRUApplication struct {
	cache *lru.Cache
	LRUController
}

type ARCApplication struct {
	cache *lru.ARCCache
	ARCController
}

type LRUController interface {
	AddCache(c *gin.Context)
	DelCache(c *gin.Context)
	GetCache(c *gin.Context)
}

type ARCController LRUController

var app *Application

func NewApplication() {
	app = &Application{}
	app.lruApp.cache, _ = lru.New(10000)
	app.arcApp.cache, _ = lru.NewARC(10000)

	g.Run(g.Config{
		ListenAddr: "localhost:8080",
		Handler: app.router(),
	})
}

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
		}
	}
	return router
}

func (lru *LRUApplication) AddCache(c *gin.Context) {
	if data, err := ioutil.ReadAll(c.Request.Body); err != nil {
		c.AbortWithError(500, err)
		return
	} else {
		key := c.Param("key")
		if key == "" {
			c.AbortWithError(500, errors.New("invalid key"))
			return
		}

		lru.cache.Add(key, data)
		c.JSON(200, string(data))
	}
}

func (lru *LRUApplication) GetCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
	}

	data, ok := lru.cache.Get(key)
	if !ok {
		c.AbortWithError(500, errors.New("no lru cache with this key: " + key))
		return
	}

	c.JSON(200, string(data.([]byte)))
}

func (lru *LRUApplication) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
	}

	lru.cache.Remove(key)

	c.Status(204)
}

func (arc *ARCApplication) AddCache(c *gin.Context) {
	if data, err := ioutil.ReadAll(c.Request.Body); err != nil {
		c.AbortWithError(500, err)
		return
	} else {
		key := c.Param("key")
		if key == "" {
			c.AbortWithError(500, errors.New("invalid key"))
			return
		}

		arc.cache.Add(key, data)
		c.JSON(200, string(data))
	}
}

func (arc *ARCApplication) GetCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
	}

	data, ok := arc.cache.Get(key)
	if !ok {
		c.AbortWithError(500, errors.New("no lru cache with this key: " + key))
		return
	}

	c.JSON(200, string(data.([]byte)))
}

func (arc *ARCApplication) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
	}

	arc.cache.Remove(key)

	c.Status(204)
}
