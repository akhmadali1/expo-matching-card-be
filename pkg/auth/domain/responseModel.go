package auth_domain

type CreateUserRequest struct {
	Error       int     `json:"error"`
	Time        int64   `json:"time"`
	Username    string  `json:"username"`
	Email       string  `json:"email"`
	Score       float64 `json:"score"`
	TimeExpired int64   `json:"time_expired"`
}

type ReqBody struct {
	Credentials string `json:"credentials"`
}
