package auth

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Detail string `json:"detail"`
}

type LogoutResponse struct {
	Detail string `json:"detail"`
}
