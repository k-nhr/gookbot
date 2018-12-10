package main

import (
	"github.com/nlopes/slack"
)

const (
	TestChannel           = "GCVB75B5H"
	AsakaiChannel         = "GCC72RSUQ"
	YukaiChannel          = "GCCMRML92"
	FieldColor1           = "#99efdf"
	FieldColor2           = "#99efdf"
	FieldColor3           = "#e6a1ed"
	FieldColor4           = "#ceeda0"
	YesterdaysTaskMessage = "昨日は何をされましたか？"
	TodaysTaskMessage     = "今日は何をしますか？"
	TroublesMessage       = "進捗を妨げるものは何ですか？"
	HolidaysMessage       = "明日のお休み、直行、帰社などあれば教えてください。"
	DoYourBestMessage     = "Awesome! Have a great day :120:"
	GoodMessage           = "今日はなにかいいことはありましたか？(Good)"
	KeepMessage           = "今日やったことで継続したいことはありますか？(Keep)"
	ProblemsMessage       = "今日あったことで問題だと思う事はありますか？(Problem)"
	TryMessage            = "これから取り組みたいことはありますか？(Try)"
	ThanksMessage         = "Thank you! Have fun :120:"
)

func (b *bot) asakaiReply(m *dailyStandupInfo, c string) {
	switch m.status {
	case TODAY:
		b.todaysTask(c, m)
	case TROUBLE:
		b.troubles(c, m)
	case HOLIDAY:
		b.holiday(c, m)
	case DOYOURBEST:
		b.doYourBest(*m)
		m.status = 0
	}
}

func (b *bot) yesterdaysTask(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: YesterdaysTaskMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question1 = YesterdaysTaskMessage
}

func (b *bot) todaysTask(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: TodaysTaskMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question2 = TodaysTaskMessage
}

func (b *bot) troubles(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: TroublesMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question3 = TroublesMessage
}

func (b *bot) holiday(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: HolidaysMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question4 = HolidaysMessage
}

func (b *bot) doYourBest(m dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: DoYourBestMessage,
	}}
	if err := b.postMessage(BotName, m.Id, attachments); err != nil {
		return
	}
	if err := b.postMessage(m.Name, AsakaiChannel, makeReport(m)); err != nil {
		return
	}
}

func (b *bot) yukaiReply(m *dailyStandupInfo, c string) {
	switch m.status {
	case KEEP:
		b.keep(c, m)
	case PROBLEM:
		b.problem(c, m)
	case TRY:
		b.try(c, m)
	case THANKS:
		b.thanks(*m)
		m.status = 0
	}
}

func (b *bot) good(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: GoodMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question1 = GoodMessage
}

func (b *bot) keep(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: KeepMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question2 = KeepMessage
}

func (b *bot) problem(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: ProblemsMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question3 = ProblemsMessage
}

func (b *bot) try(channel string, d *dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: TryMessage,
	}}
	err := b.postMessage(BotName, channel, attachments)
	if err != nil {
		return
	}
	d.question4 = TryMessage
}

func (b *bot) thanks(m dailyStandupInfo) {
	attachments := []slack.Attachment{slack.Attachment{
		Pretext: ThanksMessage,
	}}
	if err := b.postMessage(BotName, m.Id, attachments); err != nil {
		return
	}
	if err := b.postMessage(m.Name, YukaiChannel, makeReport(m)); err != nil {
		return
	}
}

func makeReport(m dailyStandupInfo) []slack.Attachment {
	attachments := make([]slack.Attachment, 0)

	if m.answer1 != "" {
		Field1 := []slack.AttachmentField{slack.AttachmentField{
			Title: m.question1,
			Value: m.answer1,
		}}
		attachments = append(attachments, slack.Attachment{
			Pretext: "*" + m.Name + "*" + " posted an update for ＊hulftiot-asakai＊",
			Fields:  Field1,
			Color:   FieldColor1,
		})
	}

	if m.answer2 != "" {
		Field2 := []slack.AttachmentField{slack.AttachmentField{
			Title: m.question2,
			Value: m.answer2,
		}}
		attachments = append(attachments, slack.Attachment{
			Fields: Field2,
			Color:  FieldColor2,
		})
	}

	if m.answer3 != "" {
		Field3 := []slack.AttachmentField{slack.AttachmentField{
			Title: m.question3,
			Value: m.answer3,
		}}
		attachments = append(attachments, slack.Attachment{
			Fields: Field3,
			Color:  FieldColor3,
		})
	}

	if m.answer4 != "" {
		Field4 := []slack.AttachmentField{slack.AttachmentField{
			Title: m.question4,
			Value: m.answer4,
		}}
		attachments = append(attachments, slack.Attachment{
			Fields: Field4,
			Color:  FieldColor4,
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
