package filter_usecase

import (
	filter_domain "match_card/pkg/filter/domain"
	filter_db "match_card/pkg/filter/repository/db"
)

func GetAllUsecase() (returnDB []filter_domain.ScoreResponse, statusResponse bool, msg string) {

	data, status, message := filter_db.GetAllScore()

	if status {
		return data, status, message
	} else {
		return data, status, msg
	}

}
