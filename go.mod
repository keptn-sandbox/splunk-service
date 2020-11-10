module example.com/splunk-service

go 1.13

replace example.com/splunk-service/client => ./splunk-service/client

require (
	github.com/cloudevents/sdk-go v1.1.2
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/keptn/go-utils v0.6.3-0.20200828090559-595b01d0f2a6
	github.com/onsi/ginkgo v1.12.0 // indirect
	github.com/onsi/gomega v1.9.0 // indirect
	github.com/sirupsen/logrus v1.2.0
	golang.org/x/sys v0.0.0-20200124204421-9fbb57f87de9 // indirect
	gopkg.in/yaml.v2 v2.2.8
)
