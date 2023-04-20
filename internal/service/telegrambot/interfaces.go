package telegrambot

import "net/http"

type ITelegramBot interface {
	SendNotification(message any) (*http.Response, error)
}
