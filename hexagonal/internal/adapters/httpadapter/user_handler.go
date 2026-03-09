package httpadapter

import (
	"encoding/json"
	"net/http"
)

type UserService interface {
	CreateUser(name, email string) error
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	err := h.service.CreateUser(req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
