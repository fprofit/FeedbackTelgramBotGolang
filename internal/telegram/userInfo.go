package telegram

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
)

var (
	DBFileName      = "userInfo.txt"
	mapIdUserInfo   = make(map[int]User)
	dbUserInfoMutex sync.RWMutex
)

func (m Message) getUserInfo(ID int) {
	var botMessage BotSendMessage
	botMessage.ChatID = m.Chat.ID
	us := GetUserInfoToMap(ID)
	botMessage.Text = fmt.Sprintf("<a href=\"tg://user?id=%d\">%s %s</a>\n\n@%s\n\n%s", us.ID, us.FirstName, us.LastName, us.Username, us.LangCode) //"<a href=\"tg://user?id=" + strconv.Itoa(us.ID) + "\">" + us.FirstName + " " + us.LastName + "</a>\n\n@" + us.Username + "\n\n" + us.LangCode
	botMessage.ParseMode = "HTML"
	buf, _ := json.Marshal(botMessage)

	PostRequestGetResponse("sendMessage", buf)
}

func AddInMapUserInfo(ID int, userInfo User) {
	dbUserInfoMutex.Lock()
	mapIdUserInfo[ID] = userInfo
	dbUserInfoMutex.Unlock()
	writeDBmapUser()
}

func GetUserInfoToMap(ID int) (user User) {
	dbUserInfoMutex.RLock()
	defer dbUserInfoMutex.RUnlock()
	if _, ok := mapIdUserInfo[ID]; ok {
		user = mapIdUserInfo[ID]
	}
	return
}

func ReadDBmapUser() {
	dbUserInfoMutex.Lock()
	defer dbUserInfoMutex.Unlock()
	_, err := os.Stat(DBFileName)
	if err != nil {
		fCreat, err := os.Create(DBFileName)
		if err != nil {
			logger.LogToFile(err)
			return
		}
		fCreat.Close()
		return
	}
	file, err := os.ReadFile(DBFileName)
	if err != nil {
		logger.LogToFile(err)
		return
	}
	errJS := json.Unmarshal(file, &mapIdUserInfo)
	if errJS != nil {
		logger.LogToFile(errJS)
	}
}

func writeDBmapUser() {
	dbUserInfoMutex.Lock()
	defer dbUserInfoMutex.Unlock()
	file, err := os.OpenFile(DBFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer file.Close()
	if err != nil {
		logger.LogToFile(err)
		return
	}
	buf, err := json.MarshalIndent(mapIdUserInfo, "", "\t")
	if err != nil {
		logger.LogToFile(err)
	}
	if _, err = file.Write(buf); err != nil {
		logger.LogToFile(err)
	}
}
