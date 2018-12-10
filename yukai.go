package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/nlopes/slack"
)

const cfgFile = "/opt/slack-bot/cfg.json"

func main() {
	token := os.Getenv("SLACK_API_TOKEN")

	list, err := readConf()

	if err != nil {
		panic(err)
	}
	bot := newBot(token)
	go bot.rtm.ManageConnection()

	var connected bool
	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				BotID = ev.Info.User.ID
				BotName = ev.Info.User.Name
				connected = true
			}
		}
		if connected {
			break
		}
	}

	bot.yukai(list)
	/*
		gocron.Every(1).Monday().At("09:15").Do(bot.asakai, members, token)
		gocron.Every(1).Tuesday().At("09:15").Do(bot.asakai, members, token)
		gocron.Every(1).Wednesday().At("09:15").Do(bot.asakai, members, token)
		gocron.Every(1).Thursday().At("09:15").Do(bot.asakai, members, token)
		gocron.Every(1).Friday().At("09:15").Do(bot.asakai, members, token)
		<-gocron.Start()
	*/
	return
}

func readConf() ([]dailyStandupInfo, error) {
	j, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	var list []dailyStandupInfo
	if err = json.Unmarshal(j, &list); err != nil {
		return nil, err
	}
	return list, nil
}
