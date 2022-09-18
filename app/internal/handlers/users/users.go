package users

import (
	"github.com/gorilla/mux"
	"net/http"
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
}

func (h *Handler) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.users.GetListUsers(r.Context())
	if err != nil {
		h.ResponseErrorJson(w, "", http.StatusBadRequest)
	}
	h.ResponseJson(w, users, 200)
}
