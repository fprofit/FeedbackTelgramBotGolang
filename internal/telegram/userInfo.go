package telegram

import (
	"fmt"
	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
	"strings"
)

func stringToHtml(s string) string {
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
	)
	return replacer.Replace(s)
}

func (tg *TG) GetUserFullInfo(userID int64) {
	chatInfo, err := tg.FuncGetUserInfo(userID)
	if err != nil {
		logger.LogToFile(err)
		return
	}

	var sb strings.Builder
	sb.WriteString("<b>User Info</b>\n\n")

	// Имя + Фамилия
	if chatInfo.FirstName != nil || chatInfo.LastName != nil {
		sb.WriteString("<b>Name:</b> ")
		if chatInfo.FirstName != nil {
			sb.WriteString(stringToHtml(*chatInfo.FirstName) + " ")
		}
		if chatInfo.LastName != nil {
			sb.WriteString(stringToHtml(*chatInfo.LastName))
		}
		sb.WriteString("\n")
	}

	// Username
	if chatInfo.Username != nil {
		sb.WriteString(fmt.Sprintf("<b>Username:</b> @%s\n", stringToHtml(*chatInfo.Username)))
	}

	// ID
	sb.WriteString(fmt.Sprintf("<b>ID:</b> %d\n", chatInfo.ID))

	// Тип аккаунта (бот, юзер, канал и т.д.)
	sb.WriteString(fmt.Sprintf("<b>Type:</b> %s\n", stringToHtml(chatInfo.Type)))

	// Язык
	if chatInfo.LanguageCode != "" {
		sb.WriteString(fmt.Sprintf("<b>Language:</b> %s\n", stringToHtml(chatInfo.LanguageCode)))
	}

	// Премиум
	if chatInfo.Prmium {
		sb.WriteString("💎 <b>Premium User</b>\n")
	}

	// Био
	if chatInfo.Bio != nil {
		sb.WriteString(fmt.Sprintf("<b>Bio:</b> %s\n", stringToHtml(*chatInfo.Bio)))
	}

	// Описание
	if chatInfo.Description != nil {
		sb.WriteString(fmt.Sprintf("<b>Description:</b> %s\n", stringToHtml(*chatInfo.Description)))
	}

	// Дата рождения
	if chatInfo.Birthdate != nil {
		bd := chatInfo.Birthdate
		sb.WriteString(fmt.Sprintf("<b>Birthdate:</b> %02d-%02d-%04d\n", bd.Day, bd.Month, bd.Year))
	}

	// Бизнес профиль
	if chatInfo.BusinessIntro != nil {
		bi := chatInfo.BusinessIntro
		sb.WriteString("<b>Business Intro:</b>\n")
		if bi.Title != nil {
			sb.WriteString(fmt.Sprintf("- Title: %s\n", stringToHtml(*bi.Title)))
		}
		if bi.Message != nil {
			sb.WriteString(fmt.Sprintf("- Message: %s\n", stringToHtml(*bi.Message)))
		}
	}

	// Локация бизнеса
	if chatInfo.BusinessLocation != nil {
		bl := chatInfo.BusinessLocation
		sb.WriteString(fmt.Sprintf("<b>Business Address:</b> %s\n", stringToHtml(bl.Address)))
		if bl.Location != nil {
			sb.WriteString(fmt.Sprintf("Location: %.6f, %.6f\n", bl.Location.Latitude, bl.Location.Longitude))
		}
	}

	// Приватности
	if chatInfo.HasPrivateForwards != nil && *chatInfo.HasPrivateForwards {
		sb.WriteString("🔒 Private forwards enabled\n")
	} else {
		// Линк на профиль
		sb.WriteString(fmt.Sprintf("<a href=\"tg://user?id=%d\">Open Profile</a>\n", chatInfo.ID))
	}
	if chatInfo.HasProtectedContent != nil && *chatInfo.HasProtectedContent {
		sb.WriteString("🚫 Protected content enabled\n")
	}
	if chatInfo.HasVisibleHistory != nil && *chatInfo.HasVisibleHistory {
		sb.WriteString("📜 Visible chat history enabled\n")
	}

	// Медиа настройки
	if chatInfo.CanSendPaidMedia != nil && !*chatInfo.CanSendPaidMedia {
		sb.WriteString("💸 Paid media not allowed\n")
	}
	if chatInfo.SlowModeDelay != nil {
		sb.WriteString(fmt.Sprintf("⌛ Slow Mode Delay: %d sec\n", *chatInfo.SlowModeDelay))
	}
	if chatInfo.MessageAutoDeleteTime != nil {
		sb.WriteString(fmt.Sprintf("🗑 Auto Delete Time: %d sec\n", *chatInfo.MessageAutoDeleteTime))
	}

	// Бусты
	if chatInfo.UnrestrictBoostCount != nil {
		sb.WriteString(fmt.Sprintf("🚀 Boost Count: %d\n", *chatInfo.UnrestrictBoostCount))
	}

	// --- Отправка ---
	var botMessage BotSendMessage
	botMessage.ChatID = tg.AdmID
	botMessage.Text = sb.String()
	botMessage.ParseMode = "HTML"

	_, err = tg.PostRequestGetResponse("sendMessage", botMessage)
	if err != nil {
		logger.LogToFile(err)
	}

}
