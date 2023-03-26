package jwt_test

import (
	"strconv"
	"testing"

	"ecommerce-evermos-projects/internal/utils/jwt"

	goJWT "github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

var id = 1234567890
var email = "example@example.com"
var secret = []byte("secretKey")

func TestCreateJWT(t *testing.T) {
	// call function CreateJWT
	tokenString, err := jwt.CreateJWT(id, email, secret)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := goJWT.ParseWithClaims(tokenString, &jwt.CustomClaims{}, func(token *goJWT.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	// assert
	assert.NoError(t, err)
	claims, ok := token.Claims.(*jwt.CustomClaims)
	assert.True(t, ok)
	assert.True(t, token.Valid)
	actID, _ := strconv.Atoi(claims.ID)
	assert.Equal(t, id, actID)
	assert.Equal(t, email, claims.Email)
	t.Log("token", tokenString)
}

func TestValidateToken(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJldmVybW9zLWVjb21tZXJjZSIsImV4cCI6OTk5OTk5OTk5OTksImlhdCI6MTY3OTg0NTcxMiwianRpIjoiMTIzNDU2Nzg5MCIsImVtYWlsIjoiZXhhbXBsZUBleGFtcGxlLmNvbSJ9.rxpNdwHkBdkLuziEix4HGAEj6mVY2ITI9_FKmJt2Jt4"

	jwtToken, err := jwt.ValidateToken(token, secret)
	assert.NoError(t, err)
	decodedToken, ok := jwtToken.Claims.(*jwt.CustomClaims)
	assert.True(t, ok)
	assert.Equal(t, email, decodedToken.Email)

	actID, _ := strconv.Atoi(decodedToken.ID)
	assert.Equal(t, id, actID)

}
