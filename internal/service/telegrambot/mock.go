package telegrambot

import (
	"encoding/json"
	"net/http"

	"smart-door/pkg/logging"

	"go.uber.org/zap"
)

type Mock struct {
	BaseURL string
	logger  logging.Logger
}

func NewTelegramBotMock(baseURL string, logger logging.Logger) *Mock {
	return &Mock{BaseURL: baseURL, logger: logger}
}

func (bot *Mock) SendNotification(message any) (*http.Response, error) {
	jsonBody, errMarshal := json.Marshal(message)
	if errMarshal != nil {
		return nil, errMarshal
	}

	bot.logger.Info("send message to telegram bot", zap.String("message", string(jsonBody)))
	return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		},
		nil
}
