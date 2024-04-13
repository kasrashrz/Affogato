package jwt

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/kasrashrz/Affogato/domain/dto_only"
	"time"
)

func GenerateJWT(email, secret string) (tokenString string, err error) {
	expirationTime := time.Now().Add(1000000 * time.Hour)
	claims := &dto_only.JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString([]byte("affogato"))
	return
}

func ValidateToken(signedToken, secret string) (err error) {
	token, err := jwt.ParseWithClaims(signedToken, &dto_only.JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("affogato"), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*dto_only.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	fmt.Printf("\n token data %v %v \n", claims.Email, claims.StandardClaims.ExpiresAt)
	return
}

func GetTokenData(signedToken, secret string) (email string) {
	token, err := jwt.ParseWithClaims(signedToken, &dto_only.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*dto_only.JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}
	return claims.Email
}
