package auth

import "encoding/json"

type ErrorResponse struct {
	Message string          `json:"message"`
	Details json.RawMessage `json:"details"`
}

type JSONMarshal interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

func (errorResponse *ErrorResponse) Marshal() []byte {
	marshal, err := json.Marshal(errorResponse)
	if err != nil {
		return nil
	}
	return marshal
}
