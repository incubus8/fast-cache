package main

import (
	"github.com/incubus8/fast-cache/rest"
	"runtime"
)

func main() {

	// Setting a higher number here allows more disk I/O calls to be scheduled, hence considerably
	// improving throughput. The extra CPU overhead is almost negligible in comparison. The
	// benchmark notes are located in badger-bench/randread.
	runtime.GOMAXPROCS(128)

	app := rest.NewApplication()
	rest.NewServer(app)
}
