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
				if m.Text == "/user_info" {
					m.getUserInfo(dbMap.GetUserID(m.ReplyToMessage.MessageID))
				} else {
					m.ReplyMessage()
				}
			} else {
				m.DelMessage()
			}
		} else if m.Chat.ID != settings.SettingsDATA.AdmID {
			if m.Text == "/start" {
				m.comandStart()
			} else {
				m.ForwMessage()
			}
		}
	} else {
		m.DelMessage()
	}
}

func (m Message) EditMessageFunc() {
	if m.Chat.ID == m.From.ID {
		switch m.From.LangCode {
		case "ru":
			SendMessage(m.Chat.ID, "Бот не может редактировать уже отправленные сообщение, отправьте сообщение заново")
		default:
			SendMessage(m.Chat.ID, "The bot can't edit already sent messages, please resend the message")
			return
		}
	}
}

func (m Message) comandStart() {
	if _, ok := settings.SettingsDATA.Text[m.From.LangCode]; ok {
		SendMessage(m.Chat.ID, settings.SettingsDATA.Text[m.From.LangCode])
	} else if _, ok := settings.SettingsDATA.Text["default"]; ok {
		SendMessage(m.Chat.ID, settings.SettingsDATA.Text["default"])
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

func SendMessage(id int, text string) {
	var botMessage BotSendMessage
	botMessage.ChatID = id
	botMessage.Text = text
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
	AddInMapUserInfo(m.Chat.ID, m.From)
	dbMap.AddInMap(postResponse.Result.MessageID, m.Chat.ID)
}
