package telegram

type GetMe struct {
    Ok     bool `json:"ok"`
    Result User `json:"result,omitempty"`
}

type User struct {
    ID        int    `json:"id`
    IsBot     bool   `json:"is_bot"`
    Username  string `json:"username,oitempty"`
    FirstName string `json:"first_name,oitempty"`
    LastName  string `json:"last_name,oitempty"`
}

type RestResponse struct {
    Ok bool `json:"ok"`
    Result []Update `json:"result"`
}

type Update struct {
    UpdateID int    `json:"update_id"`
    Message Message `json:"message"`
}

type PostResponse struct {
    Ok bool `json:"ok"`
    Result Message `json:"result"`
}

type Message struct {
    MessageID int `json:"message_id"`
    From User       `json:"from,omitempty"`
    Chat User       `json:"chat,omitempty"`
    Text string     `json:"text,omitempty"`
    ReplyToMessage ReplyToMessage `json:"reply_to_message"`

}

type ReplyToMessage struct {
    MessageID int  `json:"message_id"`
}

type BotSendMessage struct {
    ChatID int  `json:"chat_id"`
    Text string     `json:"text,omitempty"`
    FromChatID int `json:"from_chat_id,omitempty"`
    MessageID int `json:"message_id,omitempty"`
}
