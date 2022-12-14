package telegram

import (
	"encoding/json"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/dbMap"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
)

func (m Message) MessageFunc() {
	if m.Chat.ID == m.From.ID {
		if m.Chat.ID == settings.SettingsDATA.AdmID {
			if m.ReplyToMessage != nil && dbMap.GetUserID(m.ReplyToMessage.MessageID) > 0 {
				m.ReplyMessage()
			} else {
				m.DelMessage()
			}
		} else if m.Chat.ID != settings.SettingsDATA.AdmID {
			if m.Text == "/start" {
				m.SendMessageUser()
			} else {
				m.ForwMessage()
			}
		}
	} else {
		m.DelMessage()
	}
}

func (m Message) DelMessage() {
	var botMessage BotSendMessage
	botMessage.ChatID = m.Chat.ID
	botMessage.MessageID = m.MessageID
	buf, err := json.Marshal(botMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	PostRequestGetResponse("deleteMessage", buf)
}

func SendMessage(text string) {
	var botMessage BotSendMessage
	botMessage.ChatID = settings.SettingsDATA.AdmID
	botMessage.Text = text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	PostRequestGetResponse("sendMessage", buf)
}

func (m Message) SendMessageUser() {
	var botMessage BotSendMessage
	botMessage.ChatID = m.Chat.ID
	botMessage.Text = settings.SettingsDATA.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	PostRequestGetResponse("sendMessage", buf)
}

func (m Message) ReplyMessage() {
	var copyMessage BotSendMessage
	copyMessage.ChatID = dbMap.GetUserID(m.ReplyToMessage.MessageID)
	copyMessage.FromChatID = settings.SettingsDATA.AdmID
	copyMessage.MessageID = m.MessageID
	buf, err := json.Marshal(copyMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	PostRequestGetResponse("copyMessage", buf)
}
func (m Message) ForwMessage() {
	var forMessage BotSendMessage
	forMessage.ChatID = settings.SettingsDATA.AdmID
	forMessage.FromChatID = m.Chat.ID
	forMessage.MessageID = m.MessageID
	buf, err := json.Marshal(forMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	resp := PostRequestGetResponse("forwardMessage", buf)
	var postResponse PostResponse
	err = json.Unmarshal(resp, &postResponse)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	dbMap.AddInMap(postResponse.Result.MessageID, m.Chat.ID)
}
