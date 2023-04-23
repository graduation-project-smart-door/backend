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
	request, errNewRequest := http.NewRequest(http.MethodGet, service.BaseURL+SuffixOpen, nil)
	if errNewRequest != nil {
		service.logger.Error("failed create new request", zap.Error(errNewRequest))
		return nil, errNewRequest
	}

	client := &http.Client{}
	response, errPostRequest := client.Do(request)
	if errPostRequest != nil {
		service.logger.Error("failed send post request",
			zap.Error(errPostRequest), zap.String("url", service.BaseURL+SuffixOpen))
		return nil, errPostRequest
	}
	defer response.Body.Close() //nolint:nolintlint

	return response, nil
}
