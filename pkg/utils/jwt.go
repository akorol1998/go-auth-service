package utils

import (
	"errors"
	"log"
	"time"

	"github.com/akorol1998/go-auth-service/pkg/models"
	"github.com/dgrijalva/jwt-go"
)

type TokenType uint32

const (
	AccessToken TokenType = iota
	RefreshToken
)

var JwtTokensExpiration = map[TokenType]time.Duration{
	AccessToken:  time.Minute * 3,
	RefreshToken: time.Minute * 15,
}

var JwtTokenNames = map[TokenType]string{
	AccessToken:  "AccessToken",
	RefreshToken: "RefreshToken",
}

type JwtWrapper struct {
	SecretKey      string
	Issuer         string
	ExpirationMins int64
}

type JwtClaims struct {
	Id      uint64
	Email   string
	JwtType TokenType
	jwt.StandardClaims
}

// Method for generatik signed token (jwt)
func (w *JwtWrapper) GenerateToken(user models.User, tokenT TokenType, expiration time.Duration) (string, error) {
	claims := &JwtClaims{
		Id:      uint64(user.ID),
		Email:   user.Email,
		JwtType: tokenT,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(expiration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    w.Issuer,
		},
	}
	log.Printf("Generating claims - %+v", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}
	return signedToken, err
}

func (w *JwtWrapper) ValidateToken(signedToken string) (*JwtClaims, error) {
	claims := &JwtClaims{}
	log.Println("Validating token")
	_, err := jwt.ParseWithClaims(
		signedToken,
		claims,
		func(t *jwt.Token) (interface{}, error) { return []byte(w.SecretKey), nil },
	)
	if err != nil {
		return nil, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT - has expired")
	}
	return claims, err
}
