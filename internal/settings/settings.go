package settings

import (
	"encoding/json"
	"io/ioutil"
)

type Settings struct {
	BotToken string            `json:"botToken"`
	AdmID    int               `json:"adm_id"`
	Text     map[string]string `json:"text"`
}

var SettingsDATA Settings

func (s *Settings) FileReadSettings() bool {
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
