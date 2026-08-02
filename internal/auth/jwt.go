package auth

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/naresh-official/ecommerce_go/configs"
)

type AccessTokenClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`

	jwt.StandardClaims
}

type RefreshTokenClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`

	jwt.StandardClaims
}

var (
	ErrInvalidToken = jwt.ErrSignatureInvalid
)

func GenerateAccessToken(cfg *configs.JWTConfig, userID string, email string, role string) (string, error) {
	claims := AccessTokenClaims{
		UserID: userID,
		Email:  email,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.AccessTokenExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.AccessTokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(cfg *configs.JWTConfig, userID string, role string) (string, error) {
	claims := RefreshTokenClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(cfg.RefreshTokenExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.RefreshTokenSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateAccessToken(cfg *configs.JWTConfig, tokenString string) (*AccessTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AccessTokenClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(cfg.AccessTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AccessTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}

func ValidateRefreshToken(cfg *configs.JWTConfig, tokenString string) (*RefreshTokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshTokenClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(cfg.RefreshTokenSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*RefreshTokenClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, ErrInvalidToken
	}
}
