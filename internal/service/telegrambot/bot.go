package telegrambot

import (
	"bytes"
	"encoding/json"
	"net/http"

	"smart-door/pkg/logging"

	"go.uber.org/zap"
)

const (
	SuffixRecognize = "/recognize"
)

type TelegramBot struct {
	BaseURL string
	logger  logging.Logger
}

func NewTelegramBot(baseURL string, logger logging.Logger) *TelegramBot {
	return &TelegramBot{BaseURL: baseURL, logger: logger}
}

func (bot *TelegramBot) SendNotification(message any) (*http.Response, error) {
	jsonBody, errMarshal := json.Marshal(message)
	if errMarshal != nil {
		return nil, errMarshal
	}

	request, errNewRequest := http.NewRequest(http.MethodPost, bot.BaseURL+SuffixRecognize, bytes.NewReader(jsonBody))
	if errNewRequest != nil {
		bot.logger.Error("failed create new request", zap.Error(errNewRequest))
		return nil, errNewRequest
	}
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, errPostRequest := client.Do(request)
	if errPostRequest != nil {
		bot.logger.Error("failed send post request",
			zap.Error(errPostRequest), zap.String("url", bot.BaseURL+SuffixRecognize))
		return nil, errPostRequest
	}
	defer response.Body.Close() //nolint:nolintlint

	return response, nil
}
