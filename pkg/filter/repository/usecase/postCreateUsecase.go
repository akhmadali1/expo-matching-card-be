package filter_usecase

import (
	filter_domain "match_card/pkg/filter/domain"
	filter_db "match_card/pkg/filter/repository/db"
)

func PostCreateUsecase(createRequest filter_domain.CreateUserRequest) (statusResponse bool, msg string) {

	status, message := filter_db.Create(createRequest)

	if status {
		return status, message
	} else {
		return status, msg
	}

}
