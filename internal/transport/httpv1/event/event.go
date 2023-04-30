package event

import (
	"encoding/json"
	"net/http"

	"smart-door/internal/apperror"
	"smart-door/internal/dto"
	"smart-door/internal/transport/httpv1"

	"github.com/gorilla/mux"

	_ "smart-door/internal/domain"
)

type Handler struct {
	httpv1.BaseHandler
	policy Policy
}

func NewHandler(policy Policy) *Handler {
	return &Handler{policy: policy}
}

func (handler *Handler) Register(router *mux.Router) {
	router.HandleFunc("", apperror.Middleware(handler.createEvent)).Methods("POST")
}

// @Summary Creating event
// @Tags Events
// @Produce json
// @Param user body dto.CreateEvent true "event info"
// @Success 200 {object} domain.Event
// @Failure 400 {object} apperror.AppError
// @Failure 418
// @Router /api/v1/events [post]
func (handler *Handler) createEvent(writer http.ResponseWriter, request *http.Request) error {
	var event dto.CreateEvent
	if err := json.NewDecoder(request.Body).Decode(&event); err != nil {
		return apperror.ErrDecodeData
	}

	// Валидация
	if err := event.Validate(); err != nil {
		details, _ := json.Marshal(err)
		return apperror.NewAppError(err, "validation error", details)
	}

	newEvent, errCreateEvent := handler.policy.CreateEvent(request.Context(), event.ToDomain())
	if errCreateEvent != nil {
		return apperror.ErrFailedToCreate
	}

	handler.ResponseJSON(writer, newEvent, http.StatusCreated)
	return nil
}
