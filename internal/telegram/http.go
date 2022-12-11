package telegram

import (
    "fmt"
    "encoding/json"
    "net/http"
    "bytes"
    "io/ioutil"
    "time"

    "github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
    "github.com/fprofit/FeedbackTelgramBotGolang/internal/settings"
)

func FuncGetMe() (getMe GetMe) {
	resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getMe", settings.SettingsDATA.BotToken))
	if err != nil {
		logger.LogToFile(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	err = json.Unmarshal(body, &getMe)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	return
}

func GetUpdates(offset int) ([]Update) {
    for {    
        resp, err := http.Get(fmt.Sprintf("https://api.telegram.org/bot%s/getUpdates?offset=%d", settings.SettingsDATA.BotToken, offset))
        if err != nil {
            logger.LogToFile(err)
            time.Sleep(15 * time.Minute)
            continue
        } else {
            defer resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                logger.LogToFile(err)
            }
            var restResponse RestResponse
            err = json.Unmarshal(body, &restResponse)
            if err != nil {
                logger.LogToFile(err)
            }
            return restResponse.Result
        }
    }
}

func PostRequestGetResponse(method string, buf []byte) []byte{
    for{
        resp, err := http.Post(fmt.Sprintf("https://api.telegram.org/bot%s/%s", settings.SettingsDATA.BotToken, method),  "application/json", bytes.NewBuffer(buf))
        if err != nil {
            logger.LogToFile(err)
            time.Sleep(15 * time.Minute)
            continue
        } else {
            defer resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                logger.LogToFile(err)
            }
            return body
        }
    }
    
}