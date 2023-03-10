package prometheus_metrics

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const (
	envDashFileDir = "DASH_FILE_DIR"
	envPrometheusUrl = "PROMETHEUS_URL"
)

var (
	dashboardFilesDir string 
	prometheusUrl  string 
	apiClient api.Client
)



func init(){
	dashboardFilesDir = os.Getenv(envDashFileDir)
	prometheusUrl  = os.Getenv(envPrometheusUrl)

	// log.Fatalf("Please provide the value for %s and %s env variables",envDashFileDir,envPrometheusUrl)
	if dashboardFilesDir == "" {
		dashboardFilesDir = "../../sample"
	}
	if prometheusUrl == ""{
		prometheusUrl = "http://localhost:9090"
	}

	var err error

	apiClient, err = api.NewClient(api.Config{
		Address: prometheusUrl,
		Client:  &http.Client{},
	})
	if err != nil {
		log.Fatalf("got some error %s",err)
		os.Exit(1)
	}

}

func UnMarshallIntoConsumerMetrics(file string) (ConsumerMetrics, error) {

	buffer,err := loadFile(file)
	if err!=nil{
		return ConsumerMetrics{},err
	}
	var consumerMetrics ConsumerMetrics
	if err := json.Unmarshal(buffer, &consumerMetrics); err != nil {
		return ConsumerMetrics{}, err
	}
	return consumerMetrics, nil
}

func GetLabels(metric string) ([]string, error) {

	api := v1.NewAPI(apiClient)

	endTime := time.Now()
	startTime := endTime.Add(-5 * time.Minute)

	var matches []string = make([]string, 0)
	matcher := fmt.Sprintf("{%s=\"%s\"}", model.MetricNameLabel, metric)
	matches = append(matches, matcher)

	labels, _, err := api.LabelNames(context.Background(), matches, startTime, endTime)

	if err != nil {
		return nil, err
	}
	if len(labels) == 0{
		return nil,metricNotFoundError{metricName: metric}
	}

	return labels,nil
}


func loadFile(file string) ([]byte,error){
	return os.ReadFile(filepath.Join(dashboardFilesDir,file))
}