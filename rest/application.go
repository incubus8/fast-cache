package rest

import (
	"github.com/hashicorp/golang-lru"
	"github.com/karlseguin/ccache"
)

type Application struct {
	lruApp LRUApplication
	arcApp ARCApplication
	ccacheApp CCacheApplication

	Metrics
}

func NewApplication() *Application {
	app := &Application{}
	app.lruApp.cache, _ = lru.New(100)
	app.arcApp.cache, _ = lru.NewARC(100)
	app.ccacheApp.cache = ccache.New(ccache.Configure().MaxSize(100).ItemsToPrune(100))

	return app
}
