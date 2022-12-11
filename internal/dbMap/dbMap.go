package dbMap

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
)

var (
	DBFileName       = ""
	MessageID_UserID = make(map[int]int)
	dbMutex          sync.RWMutex
)

func AddInMap(messageID, userID int) {
	dbMutex.Lock()
	MessageID_UserID[messageID] = userID
	dbMutex.Unlock()
	writeDBmap()
}

func GetUserID(messageID int) int {
	dbMutex.RLock()
	defer dbMutex.RUnlock()
	return MessageID_UserID[messageID]
}

func ReadDBmap() {
	dbMutex.Lock()
	defer dbMutex.Unlock()
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
	errJS := json.Unmarshal(file, &MessageID_UserID)
	if errJS != nil {
		logger.LogToFile(errJS)
	}
}

func writeDBmap() {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	file, err := os.OpenFile(DBFileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer file.Close()
	if err != nil {
		logger.LogToFile(err)
		return
	}
	buf, err := json.MarshalIndent(MessageID_UserID, "", "\t")
	if err != nil {
		logger.LogToFile(err)
	}
	if _, err = file.Write(buf); err != nil {
		logger.LogToFile(err)
	}
}
