package telegram

import (
	"encoding/json"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/dbMap"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
)

func (up Update) DelMessage() {
	var botMessage BotSendMessage
	botMessage.ChatID = up.Message.Chat.ID
	botMessage.MessageID = up.Message.MessageID
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

func (up Update) SendMessageUser() {
	var botMessage BotSendMessage
	botMessage.ChatID = up.Message.Chat.ID
	botMessage.Text = settings.SettingsDATA.Text
	buf, err := json.Marshal(botMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	PostRequestGetResponse("sendMessage", buf)
}

func (up Update) ReplyMessage() {
	var copyMessage BotSendMessage
	copyMessage.ChatID = dbMap.GetUserID(up.Message.ReplyToMessage.MessageID)
	copyMessage.FromChatID = settings.SettingsDATA.AdmID
	copyMessage.MessageID = up.Message.MessageID
	buf, err := json.Marshal(copyMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	PostRequestGetResponse("copyMessage", buf)
}
func (up Update) ForwMessage() {
	var forMessage BotSendMessage
	forMessage.ChatID = settings.SettingsDATA.AdmID
	forMessage.FromChatID = up.Message.Chat.ID
	forMessage.MessageID = up.Message.MessageID
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
	dbMap.AddInMap(postResponse.Result.MessageID, up.Message.Chat.ID)
}
