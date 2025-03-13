package telemetry

import (
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/otel/metric"
)

var OTLPExportedTraceBytes metric.Int64Counter

func ConfigureCustomMetrics(telemetry *component.TelemetrySettings) error {
	meter := telemetry.MeterProvider.Meter("otlpexportercustommetrics")

	var err error

	OTLPExportedTraceBytes, err = meter.Int64Counter("otelcol_exporter_otlp_exported_trace_bytes",
		metric.WithDescription("Trace sizes exported by the OTLP exporter"),
		metric.WithUnit("{bytes}"))

	if err != nil {
		return err
	}

	return nil
}
