package auth_usecase

import (
	auth_domain "match_card/pkg/auth/domain"
	auth_db "match_card/pkg/auth/repository/db"
)

func GetAllUsecase() (returnDB []auth_domain.ScoreResponse, statusResponse bool, msg string) {

	data, status, message := auth_db.GetAllScore()

	if status {
		return data, status, message
	} else {
		return data, status, msg
	}

}
