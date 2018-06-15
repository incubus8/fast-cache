package rest

import (
	"errors"
	"io/ioutil"
	"net/http"

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
		c.JSON(http.StatusOK, string(data))
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

	c.JSON(http.StatusOK, string(data.([]byte)))
}

func (arc *ARCApplication) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
		return
	}

	arc.cache.Remove(key)

	c.Status(http.StatusNoContent)
}
