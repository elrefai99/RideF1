package auth

type CreateUserRequest struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}
type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type TokenRequest struct {
	Token string `json:"token" binding:"required"`
}

type RefreshPayload struct {
	Uuid string `json:"uuid" binding:"required"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
