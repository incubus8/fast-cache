package rest

import (
	"errors"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karlseguin/ccache"
)

type CCacheApplication struct {
	cache *ccache.Cache
	CCacheController
}

type CCacheController LRUController

func (ccache *CCacheApplication) AddCache(c *gin.Context) {
	if data, err := ioutil.ReadAll(c.Request.Body); err != nil {
		c.AbortWithError(500, err)
		return
	} else {
		key := c.Param("key")
		if key == "" {
			c.AbortWithError(500, errors.New("invalid key"))
			return
		}

		expiryQuery := c.DefaultQuery("expiry", "10")
		duration, err := strconv.Atoi(expiryQuery)
		if err != nil {
			c.AbortWithError(400, err)
			return
		}

		ccache.cache.Set(key, data, time.Minute * time.Duration(duration))
		c.JSON(200, string(data))
	}
}

func (ccache *CCacheApplication) GetCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
		return
	}

	item := ccache.cache.Get(key)
	if item == nil {
		c.AbortWithError(500, errors.New("no CCache with this key: " + key))
		return
	}

	if item.Expired() {
		c.AbortWithError(500, errors.New("cache is expired"))
		return
	}

	c.JSON(200, string(item.Value().([]byte)))
}

func (ccache *CCacheApplication) DelCache(c *gin.Context) {
	key := c.Param("key")
	if key == "" {
		c.AbortWithError(500, errors.New("invalid key"))
		return
	}

	ccache.cache.Delete(key)

	c.Status(204)
}