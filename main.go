package main

import (
    //"fmt"
    "encoding/json"
    "net/http"
    "bytes"
    "strconv"
    "io/ioutil"
    "os"
    "time"
)


var messId_UserId = make(map[int]int)

func main(){
    for{
        if fileReadSettings(){
            break
        }else{
            time.Sleep(30 * time.Second)
        }
    }
    offset := 0
    sendMessage("Бот перезапущен")
    for {
        updates := getUpdates(offset)
        for _, update := range updates {
            offset = update.UpdateId + 1
            if update.Message.Chat.ChatId != chatIdAdm {
                update.forwMessage()
            }else if update.Message.Chat.ChatId == chatIdAdm {
                replyMessage(update.Message.ReplyToMessage.Message_id, update.Message.Message_id)
            }
        }
    }
}

func sendMessage(text string) {
    var botMessage BotSendMessage
    botMessage.ChatId = chatIdAdm
    botMessage.Text = text
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile(err)
    }
    postRequestGetResponse("/sendMessage", buf)
}
func replyMessage(replyMessId, messId int){
    var copyMessage CopyMessage
    copyMessage.Chat_id = messId_UserId[replyMessId]
    copyMessage.From_chat_id = chatIdAdm
    copyMessage.Message_id = messId
    buf, err := json.Marshal(copyMessage)
    if err != nil {
        LogToFile(err)
    }
    postRequestGetResponse("/copyMessage", buf)
}
func (up *Update) forwMessage (){

    var forMessage ForMessage
    forMessage.Chat_id = chatIdAdm
    forMessage.From_chat_id = up.Message.Chat.ChatId
    forMessage.Message_id = up.Message.Message_id
    buf, err := json.Marshal(forMessage)
    if err != nil {
        LogToFile(err)
    }
    resp := postRequestGetResponse("/forwardMessage", buf)
    var postResponse PostResponse
    err = json.Unmarshal(resp, &postResponse)
    if err != nil {
        LogToFile(err)
    }
    messId_UserId[postResponse.Result.PostMessageId] = up.Message.Chat.ChatId

    s := time.Now().Format("15:04 [02.01.06]") + " UserId: " + strconv.Itoa(up.Message.Chat.ChatId) + " Username: " + up.Message.Chat.Username
    s = s + " First_name: " + up.Message.Chat.First_name + " Last_name: " + up.Message.Chat.Last_name
    s = s + " Text: " + up.Message.Text + "\n"
    dbTxt(s)
}

func dbTxt (s string){
    f, err := os.OpenFile("db.txt", os.O_APPEND|os.O_WRONLY, 0777)
    if err != nil {
        LogToFile(err)
    }
    defer f.Close()

    if _, err = f.WriteString(s); err != nil {
        LogToFile(err)
    }
}

func getUpdates(offset int) ([]Update) {
    for {    
        resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
        if err == nil {
            defer resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                LogToFile(err)
            }
            var restResponse RestResponse
            err = json.Unmarshal(body, &restResponse)
            if err != nil {
                LogToFile(err)
            }
            return restResponse.Result
        }
        if err != nil {
            LogToFile(err)
            time.Sleep(30 * time.Minute)
            continue
        }
    }
}

func postRequestGetResponse(method string, buf []byte) []byte{
    for{
        resp, err := http.Post(botUrl + method,  "application/json", bytes.NewBuffer(buf))
        if err == nil {
            defer resp.Body.Close()
            body, err := ioutil.ReadAll(resp.Body)
            if err != nil {
                LogToFile(err)
            }
            return body
        }
        if err != nil {
            LogToFile(err)
            time.Sleep(30 * time.Minute)
            continue
        }
    }
    
}