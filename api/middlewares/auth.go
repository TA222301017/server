package middlewares

import (
	"errors"
	"net/http"
	"server/api/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		temp := strings.Split(authHeader, " ")
		if len(temp) != 2 {
			utils.MakeResponseError(c, http.StatusUnauthorized, "authorization header must use the format Bearer TOKEN", "invalid authorization header")
			c.Abort()
			return
		}

		claims, err := utils.VerifyJWT(temp[1])

		if errors.Is(err, utils.TokenExpiredError) {
			utils.MakeResponseError(c, http.StatusUnauthorized, "token expired, please login again", err.Error())
			c.Abort()
			return
		}

		if errors.Is(err, utils.TokenMalformedError) {
			utils.MakeResponseError(c, http.StatusUnauthorized, "token is malformed", err.Error())
			c.Abort()
			return
		}

		if errors.Is(err, utils.UncheckedTokenError) {
			utils.MakeResponseError(c, http.StatusUnauthorized, "unhandled token error", err.Error())
			c.Abort()
			return
		}

		if err != nil {
			utils.MakeResponseError(c, http.StatusUnauthorized, "unauthorized", err.Error())
			c.Abort()
			return
		}

		timeDiff := claims.ExpiresAt - time.Now().Unix()
		if timeDiff > 0 && timeDiff < utils.TokenRefreshThreshold {
			if token, err := utils.MakeJWT(claims.ID); err == nil {
				c.Set("token", token)
			}
		}

		c.Set("claims", claims)
		c.Next()
	}
}
