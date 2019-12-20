package identity

import (
	"fmt"
	"strconv"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"gitlab.com/megatech/serverex/lib/log"
)

func CreateJWT(claims map[string]interface{}, secret_key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(claims))
	fmt.Println(token)
	tokenString, err := token.SignedString([]byte(secret_key))
	fmt.Println(err, tokenString)
	return tokenString, err
}

func DecodeJWT(auth string, secret_key string) (error, int64) {
	token_str := strings.TrimPrefix(auth, "Bearer ")

	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret_key), nil
	})

	if token == nil || err != nil {
		log.Error(err)
		log.Error(token_str)
		return err, 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Warn(claims)
		uid, err := strconv.ParseInt(claims["userId"].(string), 10, 64)
		if err != nil {
			log.Error(err)
			log.Error(token_str)
			log.Error(claims)
			return err, 0
		}
		return nil, uid
	} else {
		log.Error(err)
		log.Error(token_str)
		log.Error(claims)
		return err, 0
	}
}

func DecodeJWTResetPassword(auth string, secret_key string) (error, int64, string) {
	token_str := strings.TrimPrefix(auth, "Bearer ")

	token, err := jwt.Parse(token_str, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret_key), nil
	})

	if token == nil || err != nil {
		log.Error(err)
		log.Error(token_str)
		return err, 0, ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Warn(claims)
		uid, err := strconv.ParseInt(claims["resetUserId"].(string), 10, 64)
		if err != nil {
			log.Error(err)
			log.Error(token_str)
			log.Error(claims)
			return err, 0, ""
		}

		phone, ok := claims["phone"].(string)
		if !ok {
			log.Error(err)
			log.Error(token_str)
			log.Error(claims)
			return err, 0, ""
		}

		return nil, uid, phone
	} else {
		log.Error(err)
		log.Error(token_str)
		log.Error(claims)
		return err, 0, ""
	}
}
