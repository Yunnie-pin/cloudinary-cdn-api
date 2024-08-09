package middlewares

import (
	"encoding/base64"
	"net/http"
	"strings"

	"claudinary-cdn-api/config"
	"claudinary-cdn-api/data"

	"github.com/gin-gonic/gin"
)

func SimpleAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var secret string
		loadConfig, _ := config.LoadConfig()

		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && (fields[0] == "Basic") {
			decodedByte, _ := base64.StdEncoding.DecodeString(fields[1])
			decodedString := string(decodedByte)
			secret = strings.Split(decodedString, ":")[1]
		}

		if secret != loadConfig.SecretAuth {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, data.ResponseModel{
				Response:   http.StatusUnauthorized,
				Error:      "Unauthorized",
				AppID:      "claudinary-cdn-api",
				Controller: "-",
				Action:     "-",
				Result:     nil,
			})
			return
		}

		ctx.Next()

	}
}
