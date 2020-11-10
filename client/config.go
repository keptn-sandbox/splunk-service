package client

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Configuration : structure of config object
type Configuration struct {
	LogLevel log.Level `json:"log_level" yaml:"log_level"`
	File     string    `json:"conf_file" yaml:"conf_file"`
	Payload  struct {
		Proto      string `json:"proto" yaml:"proto"`
		Host       string `json:"host" yaml:"host"`
		Port       int    `json:"port" yaml:"port"`
		Token      string `json:"token" yaml:"token"`
		Index      string `json:"index" yaml:"index"`
		SourceType string `json:"sourcetype" yaml:"sourcetype"`
		Source     string `json:"source" yaml:"source"`
		Timeout    int
		Endpoints  struct {
			Health string `json:"health" yaml:"health"`
			Raw    string `json:"raw" yaml:"raw"`
		}
	}
}

// SetConfiguration : return Splunk configuration
func SetConfiguration(f string) (Configuration, error) {
	c := Configuration{}
	filePath := ""

	filePath, err := filepath.Abs(f)
	if err != nil {
		return c, err
	}
	if c.File == "" {
		c.File = filePath
	}
	content, err := ioutil.ReadFile(c.File)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(content, &c)
	if err != nil {
		return c, err
	}

	if c.Payload.Timeout == 0 {
		c.Payload.Timeout = 5
	}
	if c.LogLevel == 0 {
		log.SetLevel(4)
		c.LogLevel = log.GetLevel()
	} else {
		log.SetLevel(c.LogLevel)
	}
	if c.Payload.Host == "" {
		return c, errors.New("Host is empty. [Check Payload]")
	}
	if c.Payload.Token == "" {
		return c, errors.New("Splunk token is empty. [Check Payload]")
	}
	if c.Payload.Port == 0 {
		c.Payload.Port = 8088
	}
	if c.Payload.Proto == "" {
		c.Payload.Proto = "https"
	}
	if c.Payload.Proto != "http" && c.Payload.Proto != "https" {
		return c, errors.New("Splunk  url protocal is unknown. [Check Payload]")
	}
	if c.Payload.Endpoints.Health == "" {
		return c, errors.New("Splunk Health API is empty. [Check Payload]")
	}
	if c.Payload.Endpoints.Raw == "" {
		return c, errors.New("Splunk API Raw is empty . [Check Payload]")
	}
	return c, nil
}
