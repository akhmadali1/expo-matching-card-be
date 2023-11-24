package auth_domain

// model dari db
import (
	"time"
)

type ScoreResponse struct {
	Id         int
	Username   string
	Error      int
	Time       int64
	Score      float64
	Difficulty int
	Createdt   time.Time
}
