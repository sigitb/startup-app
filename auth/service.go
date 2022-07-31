package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(UserId int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type JwtService struct {
}

func NewService() *JwtService {
	return &JwtService{}
}

var SERCRET_KEY = []byte("BWASTARTUP_s3rcr3t_k3Y")

func (s *JwtService) GenerateToken(UserId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = UserId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken , err := token.SignedString(SERCRET_KEY)
	if err != nil{
		return signToken, err
	}

	return signToken, nil
}

func (s *JwtService) ValidateToken(token string) (*jwt.Token, error) {
	userToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SERCRET_KEY), nil
	})
	if err != nil {
		return userToken, err
	}
	return userToken, nil
}