package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	bot "github.com/k-nhr/gookbot/bot"

	"github.com/nlopes/slack"
)

const cfgFile = "/opt/slack-bot/cfg.json"

func main() {
	token := os.Getenv("SLACK_API_TOKEN")

	members, err := readConf()

	if err != nil {
		panic(err)
	}
	bot := bot.newBot(token)
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

	bot.asakai(members)
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

func readConf() ([]member, error) {
	j, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		return nil, err
	}

	var members []member
	if err = json.Unmarshal(j, &members); err != nil {
		return nil, err
	}
	return members, nil
}
