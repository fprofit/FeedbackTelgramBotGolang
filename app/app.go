package app

import(
	// "fmt"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/telegram"
)

func StartApp() {
	offset := 0
    telegram.SendMessage("Бот перезапущен")
    for {
        updates := telegram.GetUpdates(offset)
        for _, update := range updates {
            offset = update.UpdateID + 1
        	if update.Message.Chat.ID == update.Message.From.ID {
	             if update.Message.Chat.ID == settings.SettingsDATA.AdmID {
	             	if update.Message.ReplyToMessage.MessageID > 0 {
	                	go update.ReplyMessage()
	                } else {
	                	go update.DelMessage()
	                }
	            } else if update.Message.Chat.ID != settings.SettingsDATA.AdmID {
	            	if update.Message.Text == "/start" {
	            		go update.SendMessageUser()
	            	} else {
	                	go update.ForwMessage()
	            	}
	            }
            } else {
            	go update.DelMessage()
	        }
        }
    }
}