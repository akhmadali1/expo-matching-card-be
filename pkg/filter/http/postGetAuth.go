package filter_response

import (
	filter_usecase "match_card/pkg/filter/repository/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Postfilter takes a filter JSON and store in DB.
// Postfilter             godoc
// @Tags         Auths
// @Produce      json
// @Param        Auth  body      filter_domain.CreateRequest  true  "Auth JSON"
// @Success      200   {object}  filter_domain.CreateResponse
// @Router       /filter/create [post]
func GetAllUser(context *gin.Context) {

	getData, status, message := filter_usecase.GetAllUsecase()

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
