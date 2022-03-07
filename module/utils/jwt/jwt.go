package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Module struct {}


func (m *Module) CreateToken(id string) (string, error) {
	mySigningKey := []byte("testSecret")

	type MyCustomClaims struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey);
	if  err != nil {
		return "", err
	}
	return ss, err
}


type MyCustomClaims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func (m *Module) DecodeToken(tokenString string) string{

	// sample token is expired.  override time so it parses as valid

	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	claims, ok := token.Claims.(*MyCustomClaims);

	if ok && token.Valid {
		fmt.Printf("%v %v", claims.ID, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
	return claims.ID
}
