package http

import (
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	service UserService
}

type UserService interface {
	CreateUser(name, email string) error
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
