package main

import (
	//"log"

	"encoding/json"

	"github.com/cloudevents/sdk-go/pkg/cloudevents"
	keptn "github.com/keptn/go-utils/pkg/lib"

	"fmt"
	"time"

	sc "example.com/splunk-service/client"

	log "github.com/sirupsen/logrus"
)

/**
* Here are all the handler functions for the individual event
  See https://github.com/keptn/spec/blob/0.1.3/cloudevents.md for details on the payload

  -> "sh.keptn.event.configuration.change"
  -> "sh.keptn.events.deployment-finished"
  -> "sh.keptn.events.tests-finished"
  -> "sh.keptn.event.start-evaluation"
  -> "sh.keptn.events.evaluation-done"
  -> "sh.keptn.event.problem.open"
	-> "sh.keptn.events.problem"
	-> "sh.keptn.event.action.triggered"
*/

// Handles ConfigureMonitoringEventType = "sh.keptn.event.monitoring.configure"
func HandleConfigureMonitoringEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.ConfigureMonitoringEventData) error {
	log.Printf("Handling Configure Monitoring Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles ConfigurationChangeEventType = "sh.keptn.event.configuration.change"
// TODO: add in your handler code
//
func HandleConfigurationChangeEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.ConfigurationChangeEventData) error {
	log.Printf("Handling Configuration Changed Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles DeploymentFinishedEventType = "sh.keptn.events.deployment-finished"
// TODO: add in your handler code
//
func HandleDeploymentFinishedEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.DeploymentFinishedEventData) error {
	log.Printf("Handling Deployment Finished Event: %s", incomingEvent.Context.GetID())

	// capture start time for tests
	// startTime := time.Now()

	// run tests
	// ToDo: Implement your tests here

	// Send Test Finished Event
	// return myKeptn.SendTestsFinishedEvent(&incomingEvent, "", "", startTime, "pass", nil, "splunk-service")
	return nil
}

//
// Handles TestsFinishedEventType = "sh.keptn.events.tests-finished"
// TODO: add in your handler code
//
func HandleTestsFinishedEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.TestsFinishedEventData) error {
	log.Printf("Handling Tests Finished Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles EvaluationDoneEventType = "sh.keptn.events.evaluation-done"
// TODO: add in your handler code
//
func HandleStartEvaluationEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.StartEvaluationEventData) error {
	log.Printf("Handling Start Evaluation Event: %s", incomingEvent.Context.GetID())

	return nil
}

//
// Handles DeploymentFinishedEventType = "sh.keptn.events.deployment-finished"
// TODO: add in your handler code
//
func HandleEvaluationDoneEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.EvaluationDoneEventData) error {
	log.Printf("Handling Evaluation Done Event: %s", incomingEvent.Context.GetID())
	//log.Printf(" &{}", data)

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	dataField, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	configFile := "splunk/hec.config.yaml"
	conf, err := sc.SetConfiguration(configFile)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	conf.LogLevel = log.DebugLevel
	//log.SetLevel(conf.LogLevel)
	fwdclient, err := sc.SetClient(conf)
	if err != nil {
		log.Errorf("%s", err)
		return err
	}
	//log.Printf("%s", conf.Payload.Index)
	log.Debugf("%s", string(dataField))
	msg := sc.EventRequest{
		Message: fmt.Sprintf("This message on %s\"", time.Now().String()),
		Index:   conf.Payload.Index,
		Fields: map[string]string{
			"name": fwdclient.AppName,
		},
		Data:       string(dataField),
		Source:     conf.Payload.Source,
		SourceType: conf.Payload.SourceType,
	}
	log.Debugf("message=\"%v\"", msg)
	if err := fwdclient.SubmitEvent(msg); err != nil {
		log.Errorf("%s", err)
		return err
	}

	return nil
}

//
// Handles ProblemOpenEventType = "sh.keptn.event.problem.open"
// Handles ProblemEventType = "sh.keptn.events.problem"
// TODO: add in your handler code
//
func HandleProblemEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.ProblemEventData) error {
	log.Printf("Handling Problem Event: %s", incomingEvent.Context.GetID())

	// Deprecated since Keptn 0.7.0 - use the HandleActionTriggeredEvent instead

	return nil
}

//
// Handles ActionTriggeredEventType = "sh.keptn.event.action.triggered"
// TODO: add in your handler code
//
func HandleActionTriggeredEvent(myKeptn *keptn.Keptn, incomingEvent cloudevents.Event, data *keptn.ActionTriggeredEventData) error {
	log.Printf("Handling Action Triggered Event: %s", incomingEvent.Context.GetID())

	// check if action is supported
	if data.Action.Action == "action-xyz" {
		//myKeptn.SendActionStartedEvent()

		// Implement your remediation action here

		//myKeptn.SendActionFinishedEvent()
	}
	return nil
}
