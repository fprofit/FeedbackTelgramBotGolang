package settings

import (
    "io/ioutil"
    "encoding/json"
)
type Settings struct {
	BotToken string `json:"botToken"`
	AdmID int `json:"adm_id"`
    Text string `json:"text"`
}

var SettingsDATA Settings

func (s *Settings)FileReadSettings () bool{
    buf, _ := ioutil.ReadFile("settings.txt")
    err := json.Unmarshal(buf, s)
    if err != nil {
        // LogToFile(err)
        return false
    }
    if s.BotToken == "" {
        return false
    }
    return true
}