package webserver

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/karimarttila/go/simpleserver/app/util"
	"strconv"
)

var superSecret = []byte("SuperSecret")


type SSClaim struct {
    email string `json:"email"`
    jwt.StandardClaims
}

// ************ TODO: CHECK THAT ExpiresAt is seconds or time!!!!!!!!!

func CreateJsonWebToken(userEmail string) string {
	util.LogEnter()
	expStr := util.MyConfig["json_web_token_expiration_as_seconds"]
	expiration, err := strconv.Atoi(expStr)
	if err != nil {
		util.LogError("Error converting json_web_token_expiration_as_seconds: " + expStr)
	}
	myClaim := SSClaim {
		userEmail,
		jwt.StandardClaims{
			ExpiresAt: int64(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
	tokenString, err := token.SignedString(superSecret)
	if err != nil {
		util.LogError("error signing json web token: " + err.Error())
	}
	util.LogExit()
	return tokenString
}


