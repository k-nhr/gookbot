package main

import (
	"github.com/nlopes/slack"
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
