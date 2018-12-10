package main

import (
	"github.com/nlopes/slack"
)

var (
	BotID   string
	BotName string
)

type bot struct {
	api *slack.Client
	rtm *slack.RTM
}

func newBot(token string) *bot {
	bot := new(bot)
	bot.api = slack.New(token)
	bot.rtm = bot.api.NewRTM()
	return bot
}
