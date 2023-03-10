package prometheus_metrics

// to implement set as map[string]void
type void struct{}

type MetricsInGrafana struct {
	MetricsUsed    map[string]Metric  `json:"metricsUsed"`
	OverallMetrics map[string]Metric  `json:"-"`
	Dashboards     []DashboardMetrics `json:"dashboards"`
}

type DashboardMetrics struct {
	Slug        string            `json:"slug"`
	UID         string            `json:"uid,omitempty"`
	Title       string            `json:"title"`
	Metrics     map[string]Metric `json:"metrics"`
	ParseErrors []string          `json:"parse_errors"`
}

type Metric struct {
	LabelKeys map[string]void `json:"label_keys"`
}
