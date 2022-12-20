package telegram

import (
	"encoding/json"
	"fmt"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
)

func stringToHtml(str string) string {
	var str2 string
	for _, r := range []rune(str) {
		if r == '<' {
			str2 = str2 + "&lt;"
			continue
		} else if r == '>' {
			str2 = str2 + "&gt;"
			continue
		} else if r == '&' {
			str2 = str2 + "&amp;"
			continue
		} else {
			str2 = str2 + string(r)
		}
	}
	return str2
}

func (m Message) getUserInfo(ID int) {
	chatInfo := funcGetChatInfo(ID)
	if chatInfo.Ok {
		info := chatInfo.Result
		nameTohtml := stringToHtml(fmt.Sprintf("%s %s", info.FirstName, info.LastName))
		var botMessage BotSendMessage
		botMessage.ChatID = m.Chat.ID
		botMessage.Text = fmt.Sprintf("User info:\n\n<a href=\"tg://user?id=%d\">%s</a>\n\n@%s\n__________________\nBio:\n%s", info.ID, nameTohtml, info.Username, stringToHtml(info.Bio))
		botMessage.ParseMode = "HTML"
		buf, err := json.Marshal(botMessage)
		if err != nil {
			logger.LogToFile(err)
			return
		}
		PostRequestGetResponse("sendMessage", buf)
	}
}
