package telegram

import "fmt"

type RequestBody struct {
	ChatID int64 `json:"chat_id"`
	// UserID int64 `json:"user_id"`
}

func (tg *TG) banUser(id int64) {
	var requestBody RequestBody
	requestBody.ChatID = id
	// requestBody.UserID = id
	p, _ := tg.PostRequestGetResponse("leaveChat", requestBody)
	fmt.Println(StructToString(p))
	// if err != nil {
	// 	// logger.LogToFile(err)
	// }
}
