package telemetry

import "fmt"

type Config struct {
	TraceCollector  string `mapstructure:"TELEMETRY_TRACE_COLLECTOR"`
	MetricCollector string `mapstructure:"TELEMETRY_METRIC_COLLECTOR"`
	LogCollector    string `mapstructure:"TELEMETRY_LOG_COLLECTOR"`
}

func (a Config) Info() string {
	return fmt.Sprintf("Telemetry Collectors: Traces - %s, Metrics - %s, Logs - %s", a.TraceCollector, a.MetricCollector, a.LogCollector)
}
