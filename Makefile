
PACT_BROKER_URL := http://localhost:9292
PROVIDER_NAME := prometheus
CONSUMER_NAME := grafana

CONSUMER_VERSION := latest
CONTRACT_CONTENT_PATH := sample/contract.json


sample/contract.json:
	curl -XGET  \
		${PACT_BROKER_URL}/pacts/provider/${PROVIDER_NAME}/consumer/${CONSUMER_NAME}/${CONSUMER_VERSION} | jq  '.messages[0].contents' > ${CONTRACT_CONTENT_PATH}


testAndVerify:
	export PACT_BROKER_URL=${PACT_BROKER_URL} PROVIDER_NAME=${PROVIDER_NAME} CONSUMER_NAME=${CONSUMER_NAME} CONSUMER_VERSION=${CONSUMER_VERSION} && cd pkg/prometheus_metrics && go clean -testcache && go test -v .