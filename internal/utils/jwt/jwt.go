package jwt

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Email string `json:"email"`
}

func CreateJWT(ID int, email string, secret []byte) (string, error) {
	claims := new(CustomClaims)
	claims.ID = strconv.Itoa(ID)
	claims.Email = email
	claims.Issuer = "evermos-ecommerce"
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 10))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

func ValidateToken(token string, secret []byte) (*jwt.Token, error) {
	tokenClaims := &CustomClaims{}
	tokenObj, err := jwt.ParseWithClaims(token, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		hmacSampleSecret := secret
		if _, validMethod := token.Method.(*jwt.SigningMethodHMAC); !validMethod {
			return nil, errors.New("invalid token")
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !tokenObj.Valid {
		return tokenObj, errors.New("invalid token")
	}

	return tokenObj, nil
}
