package app

import (
	"time"
	
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/telegram"
)

func StartApp() {
	offset := 0
	telegram.SendMessage("Bot start")
	for {
		updates := telegram.GetUpdates(offset)
		for _, update := range updates {
			offset = update.UpdateID + 1
			if update.Message != nil {
				go update.Message.MessageFunc()
			}

		}
		time.Sleep(421 * time.Millisecond)
	}
}
