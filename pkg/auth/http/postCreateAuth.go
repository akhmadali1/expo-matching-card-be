package auth_response

import (
	"encoding/json"
	auth_domain "match_card/pkg/auth/domain"
	auth_usecase "match_card/pkg/auth/repository/usecase"
	crypto "match_card/utils/ai"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Postauth takes a auth JSON and store in DB.
// Postauth             godoc
// @Tags         Auths
// @Produce      json
// @Param        Auth  body      auth_domain.CreateRequest  true  "Auth JSON"
// @Success      200   {object}  auth_domain.CreateResponse
// @Router       /auth/create [post]
func CreateHandler(context *gin.Context) {

	var createRequest auth_domain.ReqBody
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

	var createRequestStruct auth_domain.CreateUserRequest

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

	status, message := auth_usecase.PostCreateUsecase(createRequestStruct)

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
