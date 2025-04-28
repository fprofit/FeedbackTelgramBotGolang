package telegram

type GetMe struct {
	Ok          bool         `json:"ok"`
	Result      ChatFullInfo `json:"result,omitempty"`
	ErrCode     int          `json:"error_code"`
	Description string       `json:"description"`
}

type Birthdate struct {
	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}

type RestResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId    int64    `json:"update_id,omitempty"`
	Message     *Message `json:"message,omitempty"`
	EditMessage *Message `json:"edited_message,omitempty"`
}

type PostResponse struct {
	Ok          bool     `json:"ok"`
	Result      *Message `json:"result,omitempty"`
	ErrCode     int      `json:"error_code,omitempty"`
	Description string   `json:"description,omitempty"`
}

type Message struct {
	MessageID      int64        `json:"message_id"`
	From           ChatFullInfo `json:"from,omitempty"`
	Chat           ChatFullInfo `json:"chat,omitempty"`
	Text           string       `json:"text,omitempty"`
	ReplyToMessage *Message     `json:"reply_to_message"`
}

type BotSendMessage struct {
	ChatID              int64       `json:"chat_id"`
	Text                string      `json:"text,omitempty"`
	FromChatID          int64       `json:"from_chat_id,omitempty"`
	MessageID           int64       `json:"message_id,omitempty"`
	ParseMode           string      `json:"parse_mode,omitempty"`
	Photo               string      `json:"photo,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	Entities            []Entity    `json:"entities,omitempty"`
	Media               interface{} `json:"media,omitempty"`
	ProtectContent      bool        `json:"protect_content,omitempty"`      // Optional
	DisableNotification bool        `json:"disable_notification,omitempty"` // Optional
}

type Media struct {
	Type          string   `json:"type,omitempty"`
	Media         string   `json:"media,omitempty"`
	Caption       string   `json:"caption,omitempty"`
	CaptionEntity []Entity `json:"caption_entities,omitempty"`
	HasSpoiler    bool     `json:"has_spoiler,omitempty"`
}

type Entity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"`
}
