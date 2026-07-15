package models

type GameState struct {
	Money          int
	MoneyPerSecond int
	Buildings      map[string]Building
	Timestamp      string
}
