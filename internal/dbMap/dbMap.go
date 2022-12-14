package dbMap

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
)

var (
	DBFileName       = ""
	messageID_userID = make(map[int]int)
	dbMutex          sync.RWMutex
)

func AddInMap(messageID, userID int) {
	dbMutex.Lock()
	messageID_userID[messageID] = userID
	dbMutex.Unlock()
	writeDBmap()
}

func GetUserID(messageID int) int {
	dbMutex.RLock()
	defer dbMutex.RUnlock()
	if _, ok := messageID_userID[messageID]; !ok {
		return 0
	}
	return messageID_userID[messageID]
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
	errJS := json.Unmarshal(file, &messageID_userID)
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
	buf, err := json.MarshalIndent(messageID_userID, "", "\t")
	if err != nil {
		logger.LogToFile(err)
	}
	if _, err = file.Write(buf); err != nil {
		logger.LogToFile(err)
	}
}
