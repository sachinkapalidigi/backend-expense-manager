package accesstokenutil

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/sachinkapalidigi/backend-expense-manager/domain/users"

	"github.com/sachinkapalidigi/backend-expense-manager/logger"

	"github.com/dgrijalva/jwt-go"
	"github.com/sachinkapalidigi/backend-expense-manager/utils/errors"
)

// AccessToken : JWT Token
type AccessToken struct {
	Token   string `json:"token"`
	Expires int64  `json:"expiry"`
}

const (
	expirationTime  = 24
	signingPassword = "signing_password"
)

// GenerateToken : generate jwt token based and encode uid
func GenerateToken(userID int64) (*AccessToken, *errors.RestErr) {

	var accessToken = AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      userID,
		"expires": accessToken.Expires,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv(signingPassword)))
	if err != nil {
		logger.Error("error in signing", err)
		return nil, errors.NewInternalServerError("Couldnot authenticate")
	}

	accessToken.Token = tokenString
	return &accessToken, nil
}

// ParseToken : Get user with Id on parsing token string
func ParseToken(tokenString string) *users.User {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			logger.Info("Invalid signin method for token")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv(signingPassword)), nil
	})
	if err != nil {
		logger.Error("Error in parsing", err)
		return nil
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"], claims["expires"])
		fmt.Println(reflect.TypeOf(claims["id"]))
		if userIDF, ok := claims["id"].(float64); ok {
			/* act on str */
			var userID = int64(userIDF)
			var user = users.User{ID: userID}
			if expires, ok := claims["expires"].(float64); ok {

				var at = AccessToken{
					Expires: int64(expires),
				}
				if at.IsExpired() {
					fmt.Println("Token is expired")
					return nil
				}
				return &user
			}
			fmt.Println(user)
			return &user
		} else {
			/* not string */
			fmt.Println("Not a string: userId")
			return nil
		}
	} else {
		fmt.Println("Invalid token")
		return nil
	}
}

// IsExpired : Check if token is expired
func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	atExpirationTime := time.Unix(at.Expires, 0)
	return atExpirationTime.Before(now)
}
