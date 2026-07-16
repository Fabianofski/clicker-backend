package handlers

type LoginRequest struct {
	Email    string `json:"email" validate:"required" example:"johndoe@example.com"`
	Password string `json:"password" validate:"required" example:"password"`
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required" example:"johndoe@example.com"`
	Password string `json:"password" validate:"required" example:"password"`
	Username string `json:"username" validate:"required" example:"username"`
}

type LoginResponse struct {
	Token string `json:"token" validate:"required" example:"token"`
}
