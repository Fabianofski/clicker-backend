package handlers

import (
	"net/http"

	_ "f4b1.dev/clicker-backend/internal/models"
	"f4b1.dev/clicker-backend/internal/service"
)

type LeaderboardHandler struct {
	service *service.UserService
}

func NewLeaderboardHandler() *LeaderboardHandler {
	return &LeaderboardHandler{}
}

// @Summary		State
// @Tags			Leaderboard
// @Description	Get current Leaderboard Ranking
// @Success		200	{string}	{models.LeaderboardState}	"ok"
// @Router			/leaderboard [get]
func (g *LeaderboardHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	return
}
