package rest

import (
	g "github.com/incubus8/go/pkg/gin"
)

func NewServer(app *Application) {
	g.Run(g.Config{
		ListenAddr: "localhost:8080",
		Handler: app.router(),
	})
}
