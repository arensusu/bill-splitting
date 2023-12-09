package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) *JWTMaker {
	return &JWTMaker{secretKey: secretKey}
}

type JWTPayload struct {
	UserID int64 `json:"userId"`
	jwt.RegisteredClaims
}

func (jwtMaker *JWTMaker) CreateToken(userID int64, duration time.Duration) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &JWTPayload{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        uuid.New().String(),
		},
	})

	return jwtToken.SignedString([]byte(jwtMaker.secretKey))
}

func (jwtMaker *JWTMaker) VerifyToken(token string) (*JWTPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, jwt.ErrSignatureInvalid
		}

		return []byte(jwtMaker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &JWTPayload{}, keyFunc)
	if err != nil {
		return nil, err
	}

	payload, ok := jwtToken.Claims.(*JWTPayload)
	if !ok {
		return nil, jwt.ErrInvalidType
	}

	return payload, nil
}
