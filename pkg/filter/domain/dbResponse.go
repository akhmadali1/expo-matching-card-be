package filter_domain

// model dari db
import (
	"time"
)

type ScoreResponse struct {
	Id       int
	Username string
	Time     int64
	Score    float64
	Createdt time.Time
}
