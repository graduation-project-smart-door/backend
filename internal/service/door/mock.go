package door

import (
	"net/http"

	"smart-door/pkg/logging"
)

type Mock struct {
	logger  logging.Logger
	BaseURL string
}

func NewServiceMock(logger logging.Logger, baseURL string) *Mock {
	return &Mock{BaseURL: baseURL, logger: logger}
}

func (service *Mock) Open() (*http.Response, error) {
	service.logger.Info("send message to door")
	return &http.Response{
			Status:     "200 OK",
			StatusCode: http.StatusOK,
		},
		nil
}
