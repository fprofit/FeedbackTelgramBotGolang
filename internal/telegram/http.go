package telegram

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/fprofit/FeedbackTelgramBotGolang/internal/logger"
)

func (tg *TG) FuncGetMe() error {
	resp, err := http.Get(fmt.Sprintf("%sgetMe", tg.linkAndToken))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var getMe GetMe
	err = json.Unmarshal(body, &getMe)
	if err != nil {
		return err
	}
	if getMe.Ok {
		logger.LogToFile(fmt.Errorf("Start BOT: @%s", getMe.Result.Username))
		tg.Bot = getMe.Result
	} else {
		return fmt.Errorf("Error start Bot CODE: %d, Description: %s", getMe.ErrCode, getMe.Description)
	}
	return nil
}

func (tg *TG) FuncGetUserInfo(id int64) (ChatFullInfo, error) {
	var user ChatFullInfo
	resp, err := http.Get(fmt.Sprintf("%sgetChat?chat_id=%d", tg.linkAndToken, id))
	if err != nil {
		logger.LogToFile(err)
		return user, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.LogToFile(err)
		return user, err
	}
	var getMe GetMe
	err = json.Unmarshal(body, &getMe)
	if err != nil {
		logger.LogToFile(err)
		return user, err
	}
	return getMe.Result, nil
}

func (tg *TG) GetUpdates() []Update {
	for {
		// Формируем тело запроса
		params := map[string]interface{}{
			"offset":          tg.offset,
			"timeout":         tg.timeout,
			"allowed_updates": []string{"message", "edited_message", "message_reaction"}, // здесь выбираешь нужные тебе типы
		}

		jsonData, err := json.Marshal(params)
		if err != nil {
			logger.LogToFile(err)
			time.Sleep(4 * time.Second)
			continue
		}

		// Отправляем POST-запрос
		resp, err := http.Post(tg.linkAndToken+"getUpdates", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			logger.LogToFile(err)
			time.Sleep(4 * time.Second)
			continue
		}

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.LogToFile(err)
			time.Sleep(4 * time.Second)
			continue
		}

		var restResponse RestResponse
		err = json.Unmarshal(body, &restResponse)
		if err != nil {
			logger.LogToFile(err)
			time.Sleep(4 * time.Second)
			continue
		}

		if len(restResponse.Result) == 0 {
			// logger.LogToFile(fmt.Errorf("No updates received, waiting for new updates"))
			continue
		}

		// Устанавливаем новый offset
		tg.offset = restResponse.Result[len(restResponse.Result)-1].UpdateId + 1

		return restResponse.Result
	}
}

func (tg *TG) PostRequestGetResponse(method string, b any) (PostResponse, error) {
	var result PostResponse
	var maxRetries = 3               // Максимальное количество попыток
	var retryDelay = 1 * time.Second // Время между попытками

	// Маршалинг данных
	buf, err := json.Marshal(b)
	if err != nil {
		logger.LogToFile(fmt.Errorf("Error marshaling request body: %v", err))
		return PostResponse{}, err
	}

	logger.LogToFile(fmt.Errorf("PostRequestGetResponse - Method: %s, Payload: %s", method, string(buf)))

	// Создаем HTTP-клиент с таймаутом
	client := &http.Client{
		Timeout: 5 * time.Second, // Тайм-аут для клиента
	}

	// Пытаемся несколько раз отправить запрос в случае ошибки
	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Создаем контекст с тайм-аутом на каждый запрос
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Создаем новый запрос
		req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", tg.linkAndToken, method), bytes.NewBuffer(buf))
		if err != nil {
			logger.LogToFile(fmt.Errorf("Error creating request: %v", err))
			return PostResponse{}, err
		}

		req.Header.Set("Content-Type", "application/json")

		// Выполняем запрос
		resp, err := client.Do(req)
		if err != nil {
			logger.LogToFile(fmt.Errorf("Error sending request (attempt %d/%d): %v", attempt, maxRetries, err))

			// Если попытки исчерпаны, возвращаем ошибку
			if attempt == maxRetries {
				return PostResponse{}, err
			}

			// Ждем перед следующей попыткой
			time.Sleep(retryDelay)
			continue
		}

		defer resp.Body.Close()

		// Читаем ответ
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			logger.LogToFile(fmt.Errorf("Error reading response body: %v", err))
			return PostResponse{}, err
		}

		logger.LogToFile(fmt.Errorf("Response received: %s", string(body)))

		// Декодируем ответ в структуру PostResponse
		err = json.Unmarshal(body, &result)
		if err != nil {
			logger.LogToFile(fmt.Errorf("Error unmarshalling response: %v", err))
			return PostResponse{}, err
		}
		if !result.Ok {
			var botMessage BotSendMessage
			botMessage.ChatID = tg.AdmID
			botMessage.Text = fmt.Sprintf("ErrCode: %d, Description: %s", result.ErrCode, result.Description)
			tg.PostRequestGetResponse("sendMessage", botMessage)

			return PostResponse{}, fmt.Errorf("ErrCode: %d, Description: %s", result.ErrCode, result.Description)
		}

		// Возвращаем результат, если запрос успешен
		return result, nil
	}

	// Не должно сюда доходить, но на случай, если нет успешного завершения
	return PostResponse{}, fmt.Errorf("maximum retries reached")
}
