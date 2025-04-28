package entry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/dbMap"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/telegram"
)

type Settings struct {
	BotToken string            `json:"botToken"`
	AdmID    int64             `json:"adm_id"`
	Text     map[string]string `json:"text"`
}

var SettingsDATA Settings

func FileReadSettings() error {
	buf, _ := ioutil.ReadFile("settings.json")
	err := json.Unmarshal(buf, &SettingsDATA)
	if err != nil {
		logger.LogToFile(err)
		return err
	}
	if SettingsDATA.BotToken == "" {
		return fmt.Errorf("Token null")
	}
	tg, err := telegram.NewTGBot(SettingsDATA.BotToken, 3)
	if err != nil {
		return err
	}
	tg.AdmID = SettingsDATA.AdmID
	tg.Text = SettingsDATA.Text

	dbMap.DBFileName = fmt.Sprintf("%d.txt", tg.Bot.ID)
	dbMap.ReadDBmap()

	go tg.SetComnd()
	go tg.StartApp()

	waitForShutdown()
	return nil
}

func waitForShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
}
