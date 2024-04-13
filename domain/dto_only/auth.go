package dto_only

import "github.com/golang-jwt/jwt"

//var JWTKey = []byte("L1ExZNYKh9HjUBCuk1CFa82ruFpwztX60JYJt3mwtIo=")

type JWTClaim struct {
	Email string `json:"email"`
	Code  int64  `json:"code"`
	//Pass  string `json:"pass"`
	//Iat   int64  `json:"iat,omitempty"`
	jwt.StandardClaims
}
