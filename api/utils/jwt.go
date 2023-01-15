package utils

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	Secret                  string = ""
	TokenExpirationDuration int64  = 14400
	TokenRefreshThreshold   int64  = 300
	AppName                 string = ""

	TokenExpiredError   error = errors.New("token expired")
	TokenMalformedError error = errors.New("token malformed")
	UncheckedTokenError error = errors.New("unchecked token error")
)

type CustomClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}

func init() {
	Secret = os.Getenv("API_JWT_SECRET")
	AppName = os.Getenv("APP_NAME")

	expirationDuration, err := strconv.ParseInt(os.Getenv("API_JWT_EXPIRATION"), 10, 64)
	if err != nil {
		TokenExpirationDuration = expirationDuration
	}

	refreshThreshold, err := strconv.ParseInt(os.Getenv("API_JWT_REFRESH_THRESHOLD"), 10, 64)
	if err != nil {
		TokenRefreshThreshold = refreshThreshold
	}
}

func MakeJWT(userID uint64) (string, error) {
	claims := CustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + TokenExpirationDuration,
			Issuer:    AppName,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Secret))
}

func VerifyJWT(tokenString string) (*CustomClaims, error) {
	claims := CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return &claims, nil
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, TokenMalformedError
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, TokenExpiredError
		} else {
			return nil, UncheckedTokenError
		}
	} else {
		return nil, UncheckedTokenError
	}
}
