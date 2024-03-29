package user

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
	policy      Policy
	eventBroker EventBroker
}

func NewHandler(policy Policy, eventBroker EventBroker) *Handler {
	return &Handler{policy: policy, eventBroker: eventBroker}
}

func (handler *Handler) Register(router *mux.Router) {
	router.HandleFunc("", apperror.Middleware(handler.createUser)).Methods("POST")
	router.HandleFunc("/recognize", apperror.Middleware(handler.recognizeUser)).Methods("POST")
}

// @Summary Creating a regular user
// @Tags Users
// @Produce json
// @Param user body dto.CreateUser true "user info"
// @Success 200 {object} domain.User
// @Failure 400 {object} apperror.AppError
// @Failure 418
// @Router /api/v1/users [post]
func (handler *Handler) createUser(writer http.ResponseWriter, request *http.Request) error {
	var user dto.CreateUser
	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		return apperror.ErrDecodeData
	}

	// Валидация
	if err := user.Validate(); err != nil {
		details, _ := json.Marshal(err)
		return apperror.NewAppError(err, "validation error", details)
	}

	newUser, errCreateUser := handler.policy.CreateUser(request.Context(), user.ToDomain())
	if errCreateUser != nil {
		return apperror.ErrFailedToCreate
	}

	handler.ResponseJSON(writer, newUser, http.StatusCreated)
	return nil
}

// @Summary Recognize user
// @Tags Users
// @Produce json
// @Param user body dto.RecognizeUser true "event info"
// @Success 200 {object} domain.Event
// @Failure 400 {object} apperror.AppError
// @Failure 418
// @Router /api/v1/users/recognize [post]
func (handler *Handler) recognizeUser(writer http.ResponseWriter, request *http.Request) error {
	var user dto.RecognizeUser

	if err := json.NewDecoder(request.Body).Decode(&user); err != nil {
		return apperror.ErrDecodeData
	}

	// Валидация
	if err := user.Validate(); err != nil {
		details, _ := json.Marshal(err)
		return apperror.NewAppError(err, "validation error", details)
	}

	newEvent, errCreateEvent := handler.policy.CreateEvent(request.Context(), user.ToDomain(), user.PersonID)
	if errCreateEvent != nil {
		return apperror.ErrFailedToCreate
	}

	if err := handler.eventBroker.ToMessage(*newEvent); err != nil {
		return err
	}

	handler.ResponseJSON(writer, newEvent, http.StatusCreated)
	return nil
}

// @Summary All users
// @Tags Users
// @Produce json
// @Success 200 {object} []domain.User
// @Failure 400 {object} apperror.AppError
// @Failure 418
// @Router /api/v1/users [get]
func (handler *Handler) allUsers(writer http.ResponseWriter, request *http.Request) error {
	all, errGetAllUsers := handler.policy.GetAllUsers(request.Context())
	if errGetAllUsers != nil {
		return errGetAllUsers
	}

	handler.ResponseJSON(writer, all, http.StatusOK)
	return nil
}
