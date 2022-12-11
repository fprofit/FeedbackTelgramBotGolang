package logger

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func LogToFile(error error) {
	pc, _, line, ok := runtime.Caller(1)
	os.MkdirAll("logs", 0777)
	file, err := os.OpenFile("logs/"+time.Now().Format("02.01.2006")+".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("LogToFile error os.OpenFile")
	}
	defer file.Close()
	if ok {
		_, err = file.WriteString(fmt.Sprintf("[%s] Error: %s \tline: %d %s\n", time.Now().Format("15:04"), error.Error(), line, runtime.FuncForPC(pc).Name()))
		if err != nil {
			fmt.Println("LogToFile error file.WriteString")
		}
	}
}
