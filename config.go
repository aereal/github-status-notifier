package main

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Notifications NotificationConfig
}

type NotificationConfig struct {
	Slack SlackNotificationConfig
}

type SlackNotificationConfig struct {
	WebhookURL string
	Channel    string
	Username   string
	IconEmoji  string
}

func ParseConfigFile(filename string) (*Config, error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
