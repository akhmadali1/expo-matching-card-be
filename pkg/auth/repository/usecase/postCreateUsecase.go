package auth_usecase

import (
	auth_domain "match_card/pkg/auth/domain"
	auth_db "match_card/pkg/auth/repository/db"
)

func PostCreateUsecase(createRequest auth_domain.CreateUserRequest) (statusResponse bool, msg string) {

	status, message := auth_db.Create(createRequest)

	if status {
		return status, message
	} else {
		return status, msg
	}

}
