package token

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenValid(c *gin.Context) error {
	tokenString := ExtractToken(c)
	if tokenString != os.Getenv("API_SECRET") {
		return fmt.Errorf("API NOT SAME")
	}
	return nil
}

func ExtractToken(c *gin.Context) string {
	token := c.Query("token")
	if token != "" {
		return token
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
