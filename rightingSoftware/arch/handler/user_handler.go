package handler

import (
	"encoding/json"
	"net/http"

	"github.com/emahiro/il/rightingSoftware/arch/manager"
)

type UserHandler struct {
	manager *manager.UserManager
}

func NewUserHandler(manager *manager.UserManager) *UserHandler {
	return &UserHandler{
		manager: manager,
	}
}

func (c *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	p := manager.UserRegisterParams{}
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("意図しないリクエストです。"))
		return
	}

	if err := c.manager.Register(ctx, p); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
