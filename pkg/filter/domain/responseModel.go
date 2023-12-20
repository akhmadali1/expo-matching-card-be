package filter_domain

type CreateUserRequest struct {
	Time        int64   `json:"time"`
	Username    string  `json:"username"`
	Score       float64 `json:"score"`
	TimeExpired int64   `json:"time_expired"`
}

type ReqBody struct {
	Credentials string `json:"credentials"`
}
