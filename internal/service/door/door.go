package door

import (
	"net/http"

	"smart-door/pkg/logging"

	"go.uber.org/zap"
)

const (
	SuffixOpen = "/open"
)

type Service struct {
	logger  logging.Logger
	BaseURL string
}

func NewService(logger logging.Logger, baseURL string) *Service {
	return &Service{logger: logger, BaseURL: baseURL}
}

func (service *Service) Open() (*http.Response, error) {
	response, errGetRequest := http.Get(service.BaseURL + SuffixOpen)

	if errGetRequest != nil {
		service.logger.Error("failed send get request",
			zap.Error(errGetRequest), zap.String("url", service.BaseURL+SuffixOpen))
		return nil, errGetRequest
	}

	return response, nil
}
