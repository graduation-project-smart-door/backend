package users

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"smart-door/app/internal/domain"
)

type Handler struct {
	baseHandler
	users Service
}

func NewUsersHandler(users Service) *Handler {
	return &Handler{users: users}
}

func (h *Handler) Register(router *mux.Router) {
	router.HandleFunc("/api/v1/users", h.listUsers).Methods("GET")
	router.HandleFunc("/api/v1/users", h.registrationUser).Methods("POST")
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.users.GetListUsers(r.Context())
	if err != nil {
		h.ResponseErrorJson(w, "", http.StatusBadRequest)
		return
	}
	h.ResponseJson(w, users, 200)
}

func (h *Handler) registrationUser(w http.ResponseWriter, r *http.Request) {
	var user domain.CreateUser
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		h.ResponseErrorJson(w, "wrong data", http.StatusBadRequest)
		return
	}
	err := h.users.CreateUser(r.Context(), user)
	if err != nil {
		return
	}
	h.ResponseJson(w, "created", http.StatusCreated)
}
