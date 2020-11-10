package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// HealthCheckResponse : Response variables of splunk cloud health check.
type HealthCheckResponse struct {
	Code int    `json:"code" yaml:"code"`
	Text string `json:"text" yaml:"text"`
}

// EventRequest : variables for data body
type EventRequest struct {
	Channel    string            `json:"channel,omitempty" yaml:"channel"`
	Message    string            `json:"event" yaml:"event"`
	Fields     map[string]string `json:"fields,omitempty" yaml:"fields"`
	Data       string            `json:"data" yaml:"data"`
	Host       string            `json:"host,omitempty" yaml:"host"`
	Index      string            `json:"index,omitempty" yaml:"index"`
	Source     string            `json:"source,omitempty" yaml:"source"`
	SourceType string            `json:"sourcetype,omitempty" yaml:"sourcetype"`
	Time       uint64            `json:"time,omitempty" yaml:"time"`
}

// EventResponse : response of service/raw URL
type EventResponse struct {
	Text         string `json:"text" yaml:"text"`
	Code         int    `json:"code" yaml:"code"`
	InvalidEvent int    `json:"invalid-event-number" yaml:"invalid-event-number"`
	AckID        int    `json:"ackId" yaml:"ackId"`
}

// HealthCheck : Check Splunk URL and token before HEC submit events.
func (fwdclient *Client) HealthCheck() error {
	log.Debugf("%s: url=%s", fwdclient.AppName, fwdclient.ActionUrls.Health)
	req, err := http.NewRequest("GET", fwdclient.ActionUrls.Health, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Splunk %s", fwdclient.Token))
	resp, err := fwdclient.httpclient.Do(req)
	if err != nil {
		return fmt.Errorf(" Please check splunk authorization token. %s: Health check failed: %s", fwdclient.AppName, err)
	}
	defer resp.Body.Close()
	log.Debugf("%s: status=%d %s", fwdclient.AppName, resp.StatusCode, http.StatusText(resp.StatusCode))
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s: Failed during Health check : %d %s", fwdclient.AppName, resp.StatusCode, http.StatusText(resp.StatusCode))
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s: Failed while reading health response body: %s", fwdclient.AppName, err)
	}
	healthCheckResponse := new(HealthCheckResponse)
	if err := json.Unmarshal(respBody, healthCheckResponse); err != nil {
		return fmt.Errorf("%s: health check failed: the response is not JSON but: %s", fwdclient.AppName, respBody)
	}
	log.Debugf("%s: code=%d, text=%s", fwdclient.AppName, healthCheckResponse.Code, healthCheckResponse.Text)
	return nil
}

// SubmitEvent : Post JSON message to Splunk.
func (fwdclient *Client) SubmitEvent(evt EventRequest) error {
	log.Debugf("%s: url=%s", fwdclient.AppName, fwdclient.ActionUrls.Raw)
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(evt)
	req, err := http.NewRequest("POST", fwdclient.ActionUrls.Raw, b)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Splunk %s", fwdclient.Token))
	resp, err := fwdclient.httpclient.Do(req)
	if err != nil {
		return fmt.Errorf("%s: Authorization Token is inorrect. Error: %s", fwdclient.AppName, err)
	}
	defer resp.Body.Close()
	log.Debugf("%s: status=%d %s", fwdclient.AppName, resp.StatusCode, http.StatusText(resp.StatusCode))

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s: failed: response : %s", fwdclient.AppName, err)
	}
	eventResponse := new(EventResponse)
	if err := json.Unmarshal(respBody, eventResponse); err != nil {
		return fmt.Errorf("%s: failed:  Response is not JSON formate: %s", fwdclient.AppName, respBody)
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("%s: failed: %d %s (%s)", fwdclient.AppName, resp.StatusCode, http.StatusText(resp.StatusCode), eventResponse.Text)
	}
	log.Debugf("%s: code=%d, text=%s", fwdclient.AppName, eventResponse.Code, eventResponse.Text)
	return nil
}
