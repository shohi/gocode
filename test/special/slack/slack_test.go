package slack_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/nlopes/slack"
	"github.com/pelletier/go-toml"
)

func loadSlackConfig() (token, channel string, err error) {
	confPath := "config.toml"
	tm, err := toml.LoadFile(confPath)
	if err != nil {
		return "", "", err
	}

	token = tm.Get("slack.token").(string)
	channel = tm.Get("slack.channel").(string)

	return token, channel, err
}

func checkGroup(api *slack.Client) {
	groups, err := api.GetGroups(false)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	for _, group := range groups {
		fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
	}
}

func TestSlack(t *testing.T) {
	token, channel, err := loadSlackConfig()
	if err != nil {
		t.Fatalf(err.Error())
	}

	log.Printf("channel: %v\n", channel)

	// If you set debugging, it will log all requests to the console
	// Useful when encountering issues
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	api := slack.New(token, slack.OptionDebug(true))
	checkGroup(api)
}
