package handlers

import (
	"net/http"

	_ "f4b1.dev/clicker-backend/internal/models"
	"f4b1.dev/clicker-backend/internal/service"
)

type GameHandler struct {
	service *service.UserService
}

func NewGameHandler() *GameHandler {
	return &GameHandler{}
}

// @Summary		User Game State
// @Tags			Game
// @Description	Get current Game State
// @Success		200	{string}	{models.GameState}	"ok"
// @Router			/game/state [get]
func (g *GameHandler) GetGameState(w http.ResponseWriter, r *http.Request) {
	return
}
