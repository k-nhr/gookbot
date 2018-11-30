package main

import "github.com/nlopes/slack"

const (
	AsakaiChannel         = "GCC72RSUQ"
	TestChannel           = "GCVB75B5H"
	YesterdaysTaskMessage = "昨日は何をされましたか？"
	YesterdaysTaskColor   = "#99efdf"
	TodaysTaskMessage     = "今日は何をしますか？"
	TodaysTaskColor       = "#99efdf"
	ProblemsMessage       = "進捗を妨げるものは何ですか？"
	ProblemsColor         = "#e6a1ed"
	HolidaysMessage       = "明日のお休み、直行、帰社などあれば教えてください。"
	HolidaysColor         = "#ceeda0"
	DoYourBestMessage     = "Awesome! Have a great day :120:"
)

func (b *bot) reply(m *member, c string) {
	switch m.status {
	case TODAY:
		b.todaysTask(c)
	case PROBREM:
		b.problem(c)
	case HOLIDAY:
		b.holiday(c)
	case DOYOURBEST:
		b.doYourBest(*m)
		m.status = 0
	}
}

func (b *bot) yesterdaysTask(channel string) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: YesterdaysTaskMessage,
	}}
	err := b.postMessage(botName, channel, attachments)
	if err != nil {
		return
	}
}

func (b *bot) todaysTask(channel string) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: TodaysTaskMessage,
	}}
	err := b.postMessage(botName, channel, attachments)
	if err != nil {
		return
	}
}

func (b *bot) problem(channel string) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: ProblemsMessage,
	}}
	err := b.postMessage(botName, channel, attachments)
	if err != nil {
		return
	}
}

func (b *bot) holiday(channel string) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: HolidaysMessage,
	}}
	err := b.postMessage(botName, channel, attachments)
	if err != nil {
		return
	}
}

func (b *bot) doYourBest(m member) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: DoYourBestMessage,
	}}
	if err := b.postMessage(botName, m.Id, attachments); err != nil {
		return
	}
	if err := b.postMessage(m.Name, AsakaiChannel, makeReport(m)); err != nil {
		return
	}
}

func makeReport(m member) []slack.Attachment {
	attachments := make([]slack.Attachment, 0)

	if m.ymsg != "" {
		yField := []slack.AttachmentField{slack.AttachmentField{
			Title: YesterdaysTaskMessage,
			Value: m.ymsg,
		}}
		attachments = append(attachments, slack.Attachment{
			Pretext: "*" + m.Name + "*" + " posted an update for ＊hulftiot-asakai＊",
			Fields:  yField,
			Color:   YesterdaysTaskColor,
		})
	}

	if m.tmsg != "" {
		tField := []slack.AttachmentField{slack.AttachmentField{
			Title: TodaysTaskMessage,
			Value: m.tmsg,
		}}
		attachments = append(attachments, slack.Attachment{
			Fields: tField,
			Color:  TodaysTaskColor,
		})
	}

	if m.pmsg != "" {
		pField := []slack.AttachmentField{slack.AttachmentField{
			Title: ProblemsMessage,
			Value: m.pmsg,
		}}
		attachments = append(attachments, slack.Attachment{
			Fields: pField,
			Color:  ProblemsColor,
		})
	}

	if m.hmsg != "" {
		hField := []slack.AttachmentField{slack.AttachmentField{
			Title: HolidaysMessage,
			Value: m.hmsg,
		}}
		attachments = append(attachments, slack.Attachment{
			Fields: hField,
			Color:  HolidaysColor,
		})
	}

	return attachments
}

func (b *bot) postMessage(name, channel string, attachments []slack.Attachment) error {
	params := slack.PostMessageParameters{
		Attachments: attachments,
		Username:    name,
	}
	_, _, err := b.api.PostMessage(channel, "", params)
	if err != nil {
		return err
	}
	return nil
}
