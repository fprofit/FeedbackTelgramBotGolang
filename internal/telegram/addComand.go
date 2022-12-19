package telegram

import (
	"encoding/json"
	"fmt"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
)

type MyCommands struct {
	Commands []BotCommand    `json:"commands"`
	Scope    BotCommandScope `json:"scope"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope struct {
	Type   string `json:"type"`
	ChatId int    `json:"chat_id,oitempty"`
	UserId int    `json:"user_id,oitempty"`
}

func SetComnd() {

	var botCommand BotCommand
	botCommand.Command = "user_info"
	botCommand.Description = "Reply to message"

	var botCommandScope BotCommandScope
	botCommandScope.Type = "chat"
	botCommandScope.ChatId = settings.SettingsDATA.AdmID

	var myCommands MyCommands
	myCommands.Commands = append(myCommands.Commands, botCommand)
	myCommands.Scope = botCommandScope

	buf, err := json.Marshal(myCommands)
	if err != nil {
		logger.LogToFile(err)
	}
	PostRequestGetResponse("setMyCommands", buf)
}
