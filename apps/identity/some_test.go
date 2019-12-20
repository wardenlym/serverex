package identity

import (
	"fmt"
	"io/ioutil"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/require"
)

func Test_writeconfig(t *testing.T) {
	should := require.New(t)

	b, err := toml.Marshal(default_prod_config)
	should.Nil(err)
	ioutil.WriteFile("config.prod.toml", b, 0644)
	ioutil.WriteFile("config.toml", b, 0644)
}

var hmacSampleSecret = []byte("MegaIdentitySecretKey")

func Test_createjwt(t *testing.T) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": "123",
	})
	fmt.Println(token)
	tokenString, err := token.SignedString(hmacSampleSecret)
	fmt.Println(err, tokenString)
}

func Test_ddd(t *testing.T) {
	fmt.Println(fmt.Sprintf("%06s", fmt.Sprint(21181)))
	fmt.Println(gen_code())
}
