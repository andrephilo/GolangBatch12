package models

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Gender   string `json:"gender"`
	Email    string `json:"email"`
}

type LoginResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
