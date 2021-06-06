package main
import(
    "time"
    "os"
    "fmt"
)

func LogToFile (logText string){
    file, err := os.OpenFile("logs/" + time.Now().Format("02.01.2006") + ".log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
    if err != nil {
        fmt.Println("dbTxt error os.OpenFile")
    }
    defer file.Close()

    if _, err = file.WriteString(time.Now().Format("15:04") + "\t" + logText + "\n"); err != nil {
        fmt.Println("dbTxt error file.WriteString")
    }
}