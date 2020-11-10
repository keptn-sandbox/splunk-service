// Package client : This client implements for Splunk's HTTP Event Collector (HEC).
package client

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Client : This sends messages to Splunk's
// RESTful API using HTTP/S transport.
//
type Client struct {
	httpclient *http.Client
	Token      string
	AppName    string
	ActionUrls struct {
		Health string
		Raw    string
	}
}

// SetClient : returns an instance of the Client.
func SetClient(c Configuration) (Client, error) {
	fwdclient := Client{
		AppName: "Splunk-Keptn-Client",
	}
	if err := fwdclient.SetupClient(c); err != nil {
		return fwdclient, err
	}
	log.Debugf("%s: endpoint.raw=%s", fwdclient.AppName, fwdclient.ActionUrls.Raw)
	t := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	fwdclient.httpclient = &http.Client{
		Timeout:   time.Duration(c.Payload.Timeout) * time.Second,
		Transport: t,
	}
	fwdclient.Token = c.Payload.Token
	/*if err := fwdclient.HealthCheck(); err != nil {
		return fwdclient, err
	}*/
	return fwdclient, nil
}

// SetupClient : function creates Splunk URL for endpoints
func (fwdclient *Client) SetupClient(c Configuration) error {
	fwdclient.ActionUrls.Health = fmt.Sprintf("%s://%s:%d/%s", c.Payload.Proto, c.Payload.Host, c.Payload.Port, c.Payload.Endpoints.Health)
	fwdclient.ActionUrls.Raw = fmt.Sprintf("%s://%s:%d/%s", c.Payload.Proto, c.Payload.Host, c.Payload.Port, c.Payload.Endpoints.Raw)
	return nil
}
