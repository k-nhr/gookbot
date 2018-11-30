package main

import (
	"time"

	"github.com/nlopes/slack"
)

var (
	botID   string
	botName string
)

const (
	YESTERDAY int = 1 + iota
	TODAY
	PROBREM
	HOLIDAY
	DOYOURBEST
)

type member struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	status int
	ymsg   string
	tmsg   string
	pmsg   string
	hmsg   string
}

func (bot *bot) asakai(members []member) {
	ticker := time.NewTicker(20 * time.Hour)

	for i, m := range members {
		bot.yesterdaysTask(m.Id)
		members[i].status = YESTERDAY
	}

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				exist, i := contains(members, ev.Channel)
				if !exist {
					break
				}
				if ev.Type == "message" && ev.Text != "" {
					setMessage(&members[i], ev.Text)
					bot.reply(&members[i], ev.Channel)
				}

			case *slack.DisconnectedEvent:
				return
			}
		case <-ticker.C:
			return
		}
	}
}

func contains(s []member, e string) (bool, int) {
	for i, v := range s {
		if e == v.Id {
			return true, i
		}
	}
	return false, 0
}

func setMessage(m *member, msg string) {
	if msg == "no" || msg == "-" {
		m.status++
		return
	}
	switch m.status {
	case YESTERDAY:
		m.ymsg = msg
	case TODAY:
		m.tmsg = msg
	case PROBREM:
		m.pmsg = msg
	case HOLIDAY:
		m.hmsg = msg
	}
	m.status++
}
