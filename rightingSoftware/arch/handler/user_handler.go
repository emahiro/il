package handler

import (
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
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
