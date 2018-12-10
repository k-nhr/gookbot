package main

import (
	"github.com/nlopes/slack"
)

const (
	YESTERDAY int = 1 + iota
	TODAY
	TROUBLE
	HOLIDAY
	DOYOURBEST
)

const (
	GOOD int = 1 + iota
	KEEP
	PROBLEM
	TRY
	THANKS
)

type dailyStandupInfo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	status    int
	question1 string
	question2 string
	question3 string
	question4 string
	answer1   string
	answer2   string
	answer3   string
	answer4   string
}

var Members []dailyStandupInfo

func (bot *bot) asakai(list []dailyStandupInfo) {
	for i, m := range list {
		bot.yesterdaysTask(m.Id, &list[i])
		list[i].status = YESTERDAY
	}

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				exist, i := containsMember(list, ev.Channel)
				if !exist {
					break
				}
				if ev.Type == "message" && ev.Text != "" {
					setMessage(&list[i], ev.Text)
					bot.asakaiReply(&list[i], ev.Channel)
				}

			case *slack.DisconnectedEvent:
				return
			}
		}
	}
}

func (bot *bot) yukai(list []dailyStandupInfo) {
	for i, m := range list {
		bot.good(m.Id, &list[i])
		list[i].status = YESTERDAY
	}

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				exist, i := containsMember(list, ev.Channel)
				if !exist {
					break
				}
				if ev.Type == "message" && ev.Text != "" {
					setMessage(&list[i], ev.Text)
					bot.yukaiReply(&list[i], ev.Channel)
				}

			case *slack.DisconnectedEvent:
				return
			}
		}
	}
}

func containsMember(s []dailyStandupInfo, e string) (bool, int) {
	for i, v := range s {
		if e == v.Id {
			return true, i
		}
	}
	return false, 0
}

func setMessage(m *dailyStandupInfo, msg string) {
	if msg == "no" || msg == "-" {
		m.status++
		return
	}
	switch m.status {
	case YESTERDAY:
		m.answer1 = msg
	case TODAY:
		m.answer2 = msg
	case TROUBLE:
		m.answer3 = msg
	case HOLIDAY:
		m.answer4 = msg
	}
	m.status++
}
