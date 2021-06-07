package main

import (
    "io/ioutil"
    "encoding/json"
)
type Settings struct {
	BotToken string `json:"botToken"`
	ChatIdAdm int `json:"chatIdAdm"`
}
var (
	botUrl string
	chatIdAdm int
)

func fileReadSettings () bool{
    buf, _ := ioutil.ReadFile("settings.txt")
    var settings Settings
    err := json.Unmarshal(buf, &settings)
    if err != nil {
        LogToFile("Err getUpdates json.Unmarshal")
        return false

    }
    botUrl = "https://api.telegram.org/bot" + settings.BotToken
    chatIdAdm = settings.ChatIdAdm
    return true
}