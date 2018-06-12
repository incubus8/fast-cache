package rest

import (
	"github.com/hashicorp/golang-lru"
)

type Application struct {
	lruApp LRUApplication
	arcApp ARCApplication

	Metrics
}

func NewApplication() *Application {
	app := &Application{}
	app.lruApp.cache, _ = lru.New(10000)
	app.arcApp.cache, _ = lru.NewARC(10000)

	return app
}
