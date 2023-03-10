package prometheus_metrics

import (
	"fmt"
	"os"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/assert"
)

func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "prometheus",
		LogLevel: "DEBUG",
		PactDir: "../../sample",
	}
}

var pact = createPact()

var (
	pact_broker_url string = os.Getenv("PACT_BROKER_URL")
	provider_version string = os.Getenv("PROVIDER_VERSION")
	consumer_version string = os.Getenv("CONSUMER_VERSION")
	provider_name string = os.Getenv("PROVIDER_NAME")
	consumer_name string = os.Getenv("CONSUMER_NAME")
	ci string = os.Getenv("CI")
)



func TestMetricsPact(t *testing.T) {

	pact.Setup(true)

	responseMetricsInGrafana := MetricsInGrafana{
		MetricsUsed: make(map[string]Metric),
	}

	functionMappings := dsl.MessageHandlers{
		"analyse metrics": func(m dsl.Message) (interface{}, error) {

			file := "../sample/contract.json"

			metricsInGrafana,err := UnMarshallIntoMetricsInGrafana(file)
		
			if err != nil{
				t.Errorf("got some error %s",err)
			}

		
			for metric := range metricsInGrafana.MetricsUsed {
				testName := fmt.Sprintf("running test for %s",metric)
		
				t.Run(testName, func(t *testing.T) {
					
					var responseActualLabelKeys map[string]void = make(map[string]void)

					var expectedLabelKeys []string = make([]string, 0)
					// get the label keys
					for label := range metricsInGrafana.MetricsUsed[metric].LabelKeys {
						expectedLabelKeys = append(expectedLabelKeys, label)
						responseActualLabelKeys[label] = void{}
					}
		
					actualLabelKeys, err := GetLabels(metric)
					
					if err != nil {
						t.Errorf("got some error while retrieving the labels from prometheus: %s", err)
					} else {
						// check all the expectedLabelKeys are present in actualLabelKeys
						t.Log(actualLabelKeys,expectedLabelKeys,responseActualLabelKeys)
						assert.Subset(t, actualLabelKeys, expectedLabelKeys)
						responseMetricsInGrafana.MetricsUsed[metric] = Metric{LabelKeys:responseActualLabelKeys}
					}
		
				})
			}
			return responseMetricsInGrafana, nil
		},
	}



	// Verify the Provider with local Pact Files
	pact.VerifyMessageProvider(t,dsl.VerifyMessageRequest{
		PactURLs:                   []string{fmt.Sprintf("%s/pacts/provider/%s/consumer/%s/%s",pact_broker_url,provider_name,consumer_name,consumer_version)},
		MessageHandlers:            functionMappings,
		PactLogLevel:               "INFO",		
		PublishVerificationResults: func() bool{
			return ci == "true"
		}(),
		ProviderVersion:            provider_version,
	})

	pact.Teardown()

}
