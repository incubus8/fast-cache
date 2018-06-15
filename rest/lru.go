package rest

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/golang-lru"
)

type LRUApplication struct {
	cache *lru.Cache
	LRUController
}

type LRUController interface {
	AddCache(c *gin.Context)
	DelCache(c *gin.Context)
	GetCache(c *gin.Context)
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
		c.JSON(http.StatusOK, string(data))
	}
}

func (lru *LRUApplication) GetCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
		return
	}

	data, ok := lru.cache.Get(key)
	if !ok {
		c.AbortWithError(500, errors.New("no lru cache with this key: " + key))
		return
	}

	c.JSON(http.StatusOK, string(data.([]byte)))
}

func (lru *LRUApplication) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
		return
	}

	lru.cache.Remove(key)

	c.Status(http.StatusNoContent)
}