package auth_response

import (
	auth_usecase "match_card/pkg/auth/repository/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Postauth takes a auth JSON and store in DB.
// Postauth             godoc
// @Tags         Auths
// @Produce      json
// @Param        Auth  body      auth_domain.CreateRequest  true  "Auth JSON"
// @Success      200   {object}  auth_domain.CreateResponse
// @Router       /auth/create [post]
func GetAllUser(context *gin.Context) {

	getData, status, message := auth_usecase.GetAllUsecase()

	if status {
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    getData,
			"message": message,
		})
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"data":    getData,
			"message": message,
		})
		return
	}
}
