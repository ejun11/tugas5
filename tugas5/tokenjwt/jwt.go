package tokenjwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Token struct {
	UserId string `json:"user_id"`
	Email  string `json:"email"`
}

var SECRET_KEY = "SecretKey"

func CreateToken(payload *Token) (string, error) {
	claims := jwt.MapClaims{
		"payload": payload,
		"issued":  time.Now().Add(24 * time.Hour),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokString string) (*Token, error) {
	tok, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing : %v", t.Header["tes"])
		}
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if !tok.Valid {
		return nil, errors.New("invalid token")

	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token")

	}

	issuedString := fmt.Sprintf("%v", claims["issued"])
	issued, err := time.Parse(time.RFC3339, issuedString)
	if err != nil {
		return nil, err
	}

	if time.Now().After(issued) {
		return nil, errors.New("token expired")
	}

	byteClaims, err := json.Marshal(claims["payload"])
	if err != nil {
		return nil, err
	}

	var myToken Token
	err = json.Unmarshal(byteClaims, &myToken)
	if err != nil {
		return nil, err
	}

	return &myToken, nil

}

func GeneratePassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func ValidatePassword(hash, pass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err
}
