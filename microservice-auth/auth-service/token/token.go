package token

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"golang.org/x/oauth2"
)

var privateKey *rsa.PrivateKey

func init() {
	privateKey = loadRSAPrivateKeyFromDisk("./sample_key")
}

func GenerateToken(c jwt.MapClaims) *oauth2.Token {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodRS256, c)
	claims := jwtToken.Claims.(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + oneday()
	claims["typ"] = "Bearer"

	accessToken, err := jwtToken.SignedString(privateKey)
	if err != nil {
		log.Errorf("signing failed, err:%v", err)
		return nil
	}

	token := &oauth2.Token{
		AccessToken: accessToken,
		Expiry:      time.Unix(iat+oneday(), 0),
		TokenType:   "Bearer",
	}
	return token
}

func loadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := ioutil.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func oneday() int64 {
	return cast.ToInt64(1 * 24 * 60 * 60)
}
