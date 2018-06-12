package rest

import (
	g "github.com/incubus8/go/pkg/gin"
)

func NewServer(app *Application) {
	conf := g.Config{
		ListenAddr: "localhost:8080",
		Handler: app.WrapOchttp(app),
	}

	g.Run(conf)
}
