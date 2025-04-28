package telegram

import (
	"encoding/json"
	"fmt"
)

type TG struct {
	AdmID int64
	Text  map[string]string

	linkAndToken string
	offset       int64
	timeout      int
	Bot          ChatFullInfo
}

func NewTGBot(token string, timeout int) (*TG, error) {
	tg := &TG{
		linkAndToken: fmt.Sprintf("https://api.telegram.org/bot%s/", token),
		timeout:      timeout,
	}
	if err := tg.FuncGetMe(); err != nil {
		return nil, err
	}
	return tg, nil
}

func (tg *TG) StartApp() {
	for {
		updates := tg.GetUpdates()
		for _, update := range updates {
			if update.Message != nil {
				go tg.MessageFunc(update.Message)
			} else if update.EditMessage != nil {
				go tg.EditMessageFunc(update.EditMessage)
			}
		}

	}
}

func StructToString(b any) string {
	res, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		return ""
	}
	return string(res)
}
