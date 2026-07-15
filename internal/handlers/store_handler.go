package handlers

import (
	"net/http"

	_ "f4b1.dev/clicker-backend/internal/models"
	"f4b1.dev/clicker-backend/internal/service"
)

type StoreHandler struct {
	service *service.UserService
}

func NewStoreHandler() *StoreHandler {
	return &StoreHandler{}
}

// @Summary		Purchase a specific Building
// @Tags			Store
// @Description	Get current Store State
// @Success		200	{string}	{models.StoreState}	"ok"
// @Router			/store/purchase [Post]
func (g *StoreHandler) Purchase(w http.ResponseWriter, r *http.Request) {
	return
}
