package app

import (
	"time"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/telegram"
)

func StartApp() {
	offset := 0
	telegram.SendMessage(settings.SettingsDATA.AdmID, "Bot start")
	telegram.SetComnd()
	for {
		updates := telegram.GetUpdates(offset)
		for _, update := range updates {
			offset = update.UpdateID + 1
			if update.Message != nil {
				go update.Message.MessageFunc()
			} else if update.EditMessage != nil {
				go update.EditMessage.EditMessageFunc()
			}

		}
		time.Sleep(421 * time.Millisecond)
	}
}
