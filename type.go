package main

type RestResponse struct {
    Ok bool `json:"ok"`
    Result []Update `json:"result"`
}

type Update struct {
    UpdateId int    `json:"update_id"`
    Message Message `json:"message"`
}

type Message struct {
    Message_id int `json:"message_id"`
    Chat Chat       `json:"chat"`
    Text string     `json:"text"`
    ReplyToMessage ReplyToMessage `json:"reply_to_message"`

}
type ReplyToMessage struct {
    Message_id int  `json:"message_id"`
}

type Chat struct {
    ChatId int          `json:"id"`
    Username string     `json:"username"`
    First_name string   `json:"first_name"`
    Last_name string    `json:"last_name"`
}

type ForMessage  struct {
    Chat_id int `json:"chat_id"`
    From_chat_id int `json:"from_chat_id"`
    Message_id int `json:"message_id"`
}

type CopyMessage struct {
    Chat_id int `json:"chat_id"`
    From_chat_id int `json:"from_chat_id"`
    Message_id int `json:"message_id"`

}

type PostResponse struct {
    Ok bool `json:"ok"`
    Result PostUpdate `json:"result"`
}

type PostUpdate struct {
    PostMessageId int `json:"message_id"`
}

type BotSendMessage struct {
    ChatId int  `json:"chat_id"`
    Text string     `json:"text"`
}
