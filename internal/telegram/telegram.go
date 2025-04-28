package telegram

import (
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/dbMap"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
)

func (tg *TG) MessageFunc(m *Message) {
	if m.Chat.ID == m.From.ID {
		if m.Chat.ID == tg.AdmID {
			if m.ReplyToMessage != nil && dbMap.GetUserID(m.ReplyToMessage.MessageID) > 0 {
				if m.Text == "/user_info" {
					tg.GetUserFullInfo(dbMap.GetUserID(m.ReplyToMessage.MessageID))
				} else if m.Text == "/ban" {
					tg.banUser(dbMap.GetUserID(m.ReplyToMessage.MessageID))
				} else {
					tg.ReplyMessage(m)
				}
			} else {
				tg.DelMessage(m)
			}
		} else if m.Chat.ID != tg.AdmID {
			if m.Text == "/start" {
				tg.comandStart(m)
			} else {
				tg.ForwMessage(m)
			}
		}
	} else {
		tg.DelMessage(m)
	}
}

func (tg *TG) EditMessageFunc(m *Message) {
	if m.Chat.ID == m.From.ID {
		var botMessage BotSendMessage
		botMessage.ChatID = m.Chat.ID
		switch m.From.LanguageCode {
		case "ru":
			botMessage.Text = "Бот не может редактировать уже отправленные сообщение, отправьте сообщение заново"
		default:
			botMessage.Text = "The bot can't edit already sent messages, please resend the message"
		}
		tg.PostRequestGetResponse("sendMessage", botMessage)
	}
}

func (tg *TG) comandStart(m *Message) {
	var botMessage BotSendMessage
	botMessage.ChatID = m.From.ID
	if _, ok := tg.Text[m.From.LanguageCode]; ok {
		botMessage.Text = tg.Text[m.From.LanguageCode]
	} else if _, ok := tg.Text["default"]; ok {
		botMessage.Text = tg.Text["default"]
	} else {
		return
	}
	tg.PostRequestGetResponse("sendMessage", botMessage)
}

func (tg *TG) DelMessage(m *Message) {
	var botMessage BotSendMessage
	botMessage.ChatID = m.Chat.ID
	botMessage.MessageID = m.MessageID
	tg.PostRequestGetResponse("deleteMessage", botMessage)
}

func (tg *TG) ReplyMessage(m *Message) {
	var copyMessage BotSendMessage
	copyMessage.ChatID = dbMap.GetUserID(m.ReplyToMessage.MessageID)
	copyMessage.FromChatID = tg.AdmID
	copyMessage.MessageID = m.MessageID
	_, err := tg.PostRequestGetResponse("copyMessage", copyMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
}

func (tg *TG) ForwMessage(m *Message) {
	var forMessage BotSendMessage
	forMessage.ChatID = tg.AdmID
	forMessage.FromChatID = m.Chat.ID
	forMessage.MessageID = m.MessageID
	resp, err := tg.PostRequestGetResponse("forwardMessage", forMessage)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	dbMap.AddInMap(resp.Result.MessageID, m.Chat.ID)
}
