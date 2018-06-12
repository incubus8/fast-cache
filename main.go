package main

import (
	"github.com/incubus8/fast-cache/rest"
)

func main() {
	app := rest.NewApplication()

	rest.NewServer(app)
}
