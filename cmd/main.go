package main

import (
	"fmt"

	"github.com/fprofit/FeedbackTelgramBotGolang/app"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/dbMap"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/telegram"
)

func main() {
	if settings.SettingsDATA.FileReadSettings() {
		info := telegram.FuncGetMe()
		if info.Ok {
			dbMap.DBFileName = fmt.Sprintf("%d.txt", info.Result.ID)
			dbMap.ReadDBmap()
			telegram.ReadDBmapUser()
			fmt.Println(fmt.Sprintf("Start BOT: @%s", info.Result.Username))
			app.StartApp()
		} else {
			fmt.Println("Error BOT token")
		}
	} else {
		fmt.Println("Error file setting")
	}
}
