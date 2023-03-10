package prometheus_metrics

// func TestMetrics(t *testing.T) {

// 	file := "../sample/single-metric.json"

// 	metricsInGrafana,err := UnMarshallIntoMetricsInGrafana(file)

// 	if err != nil{
// 		t.Errorf("got some error %s",err)
// 	}

// 	for metric := range metricsInGrafana.MetricsUsed {
// 		testName := fmt.Sprintf("running test for %s",metric)

// 		t.Run(testName, func(t *testing.T) {

// 			var expectedLabelKeys []string = make([]string, 0)
// 			// get the label keys
// 			for label := range metricsInGrafana.MetricsUsed[metric].LabelKeys {
// 				expectedLabelKeys = append(expectedLabelKeys, label)
// 			}

// 			actualLabelKeys, err := GetLabels(metric)

// 			if err != nil {
// 				t.Errorf("got some error while retrieving the labels from prometheus: %s", err)
// 			} else {

// 				// check all the expectedLabelKeys are present in actualLabelKeys
// 				assert.Subset(t, actualLabelKeys, expectedLabelKeys)
// 			}

// 		})
// 	}
// }
	