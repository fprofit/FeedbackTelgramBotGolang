package telegram

type MyCommands struct {
	Commands []BotCommand    `json:"commands"`
	Scope    BotCommandScope `json:"scope"`
}

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

type BotCommandScope struct {
	Type   string `json:"type"`
	ChatID int64  `json:"chat_id,oitempty"`
	UserID int64  `json:"user_id,oitempty"`
}

func (tg *TG) SetComnd() {
	var myCommands MyCommands
	myCommands.Commands = []BotCommand{{Command: "user_info", Description: "🆔 Reply to message"}} //, {Command: "ban", Description: "❌ Ban"}, {Command: "un_ban", Description: "✅ UnBan"}}
	myCommands.Scope.Type = "chat"
	myCommands.Scope.ChatID = tg.AdmID

	tg.PostRequestGetResponse("setMyCommands", myCommands)
}
