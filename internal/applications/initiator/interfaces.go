package initiator

import "net/http"

type TelegramBot interface {
	SendNotification(message any) (*http.Response, error)
}

type DoorService interface {
	Open() (*http.Response, error)
}
