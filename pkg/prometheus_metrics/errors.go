package prometheus_metrics

import "fmt"


type metricNotFoundError struct {
	metricName string 
}

func (m metricNotFoundError) Error() string{
	return fmt.Sprintf("Metric %s not found",m.metricName)
}