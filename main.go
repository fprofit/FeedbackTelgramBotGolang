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
        updates, err := getUpdates(offset)
        if err != nil {
            LogToFile("main getUpdates  err")
            time.Sleep(30 * time.Second)
        }
        for _, update := range updates {
            offset = update.UpdateId + 1
            if update.Message.Chat.ChatId != chatIdAdm {
                update.forwMessage()
            }else if update.Message.Chat.ChatId == chatIdAdm {
                replyMessage(update.Message.ReplyToMessage.Message_id, update.Message.Message_id)
            }
        }
        //LogToFile(updates)
    }
	
}
func sendMessage(text string) {
    var botMessage BotSendMessage
    botMessage.ChatId = chatIdAdm
    botMessage.Text = text
    buf, err := json.Marshal(botMessage)
    if err != nil {
        LogToFile("sendMessage error json.Marshal")
    }
    _, err = http.Post(botUrl + "/sendMessage",  "application/json", bytes.NewBuffer(buf))
    if err != nil {
        LogToFile("sendMessage error http.Post")
    }
}
func replyMessage(replyMessId, messId int){
    var copyMessage CopyMessage
    copyMessage.Chat_id = messId_UserId[replyMessId]
    copyMessage.From_chat_id = chatIdAdm
    copyMessage.Message_id = messId
    buf, err := json.Marshal(copyMessage)
    if err != nil {
        LogToFile("replyMessage error json.Marshal")
    }
    resp, err := http.Post(botUrl + "/copyMessage",  "application/json", bytes.NewBuffer(buf))
    if err != nil {
        LogToFile("replyMessage error http.Post")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        LogToFile("replyMessage error ioutil.ReadAll")
    }
    var postResponse PostResponse
    err = json.Unmarshal(body, &postResponse)
    if err != nil {
        LogToFile("replyMessage error json.Unmarshal")
    }
}
func (up *Update) forwMessage (){

    var forMessage ForMessage
    forMessage.Chat_id = chatIdAdm
    forMessage.From_chat_id = up.Message.Chat.ChatId
    forMessage.Message_id = up.Message.Message_id
    buf, err := json.Marshal(forMessage)
    if err != nil {
        LogToFile("ForwMessage error json.Marshal")
    }
    resp, err := http.Post(botUrl + "/forwardMessage",  "application/json", bytes.NewBuffer(buf))
    if err != nil {
        LogToFile("ForwMessage error http.Post")
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        LogToFile("ForwMessage error ioutil.ReadAll")
    }
    var postResponse PostResponse
    err = json.Unmarshal(body, &postResponse)
    if err != nil {
        LogToFile("ForwMessage error json.Unmarshal")
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
        LogToFile("dbTxt error os.OpenFile")
    }
    defer f.Close()

    if _, err = f.WriteString(s); err != nil {
        LogToFile("dbTxt error f.WriteString")
    }
}

func getUpdates(offset int) ([]Update, error) {
    resp, err := http.Get(botUrl + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }
    var restResponse RestResponse
    err = json.Unmarshal(body, &restResponse)
    if err != nil {
        return nil, err
    }
    return restResponse.Result, nil
}