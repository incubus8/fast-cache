package rest

import (
	"errors"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/golang-lru"
)

type ARCApplication struct {
	cache *lru.ARCCache
	ARCController
}

type ARCController LRUController

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
		return
	}

	data, ok := arc.cache.Get(key)
	if !ok {
		c.AbortWithError(500, errors.New("no ARC LRU cache with this key: " + key))
		return
	}

	c.JSON(200, string(data.([]byte)))
}

func (arc *ARCApplication) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
		return
	}

	arc.cache.Remove(key)

	c.Status(204)
}
