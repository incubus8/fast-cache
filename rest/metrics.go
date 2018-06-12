package rest

import (
	"net/http"

	"github.com/rs/zerolog/log"
	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/stats/view"
)

type Metrics struct {}

func (Metrics) WrapOchttp(app *Application) http.Handler {
	h := &ochttp.Handler{Handler: app.router()}
	if err := view.Register(ochttp.DefaultServerViews...); err != nil {
		log.Fatal().Msg("Failed to register ochttp.DefaultServerViews")
	}

	return h
}

func (Metrics) GetPrometheusExporter() *prometheus.Exporter {
	pe, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "fast-cache",
	})
	if err != nil {
		log.Fatal().Msgf("Failed to create Prometheus exporter: %v", err)
	}
	view.RegisterExporter(pe)

	return pe
}