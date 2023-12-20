package filter_response

import (
	"encoding/json"
	filter_domain "match_card/pkg/filter/domain"
	filter_usecase "match_card/pkg/filter/repository/usecase"
	crypto "match_card/utils/ai"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Postfilter takes a filter JSON and store in DB.
// Postfilter             godoc
// @Tags         Auths
// @Produce      json
// @Param        Auth  body      filter_domain.CreateRequest  true  "Auth JSON"
// @Success      200   {object}  filter_domain.CreateResponse
// @Router       /filter/create [post]
func CreateHandler(context *gin.Context) {

	var createRequest filter_domain.ReqBody
	if err := context.ShouldBindJSON(&createRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": err.Error()})
		return
	}

	if createRequest.Credentials == "" {
		context.JSON(http.StatusExpectationFailed, gin.H{"status": 417, "message": "Read credentials failed"})
		return
	}

	decrypt, err := crypto.Decrypt(createRequest.Credentials)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{"status": 417, "message": "Decrypt failed"})
		return
	}

	var createRequestStruct filter_domain.CreateUserRequest

	decryptData := json.Unmarshal([]byte(decrypt), &createRequestStruct)

	if decryptData != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{"status": 417, "message": "Decrypt Data failed"})
		return
	}

	currentTimestamp := time.Now().UTC().Unix()
	difference := currentTimestamp - createRequestStruct.TimeExpired
	MinutesInSeconds := int64(1 * 30)
	if difference > MinutesInSeconds {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "The credetials is over lifespan from the current time.",
		})
		return
	}

	status, message := filter_usecase.PostCreateUsecase(createRequestStruct)

	if status {
		context.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": message,
		})
		return
	} else {
		context.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": message,
		})
		return
	}

}
