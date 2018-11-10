package webserver

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/karimarttila/go/simpleserver/app/util"
	"strconv"
	"time"
)

// NOTE: This is an exercise. In Production we would get this e.g. from some key vault.
var superSecret = []byte("SuperSecret")

type SSClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type TokenResponse struct {
	Ready bool
	Email string `json:"email"`
}

// NOTE: Go does not have native set. We use map to simulate set.
var mySessions = make(map[string]bool)

// ************ TODO: CHECK THAT ExpiresAt is seconds or time!!!!!!!!!

func CreateJsonWebToken(userEmail string) (ret string, err error) {
	util.LogEnter()
	expStr := util.MyConfig["json_web_token_expiration_as_seconds"]
	expiration, err := strconv.Atoi(expStr)
	if err != nil {
		util.LogError("Error converting json_web_token_expiration_as_seconds: " + expStr)
	} else {
		ttl := time.Duration(expiration) * time.Second
		claimExp := time.Now().UTC().Add(ttl).Unix()

		myClaim := SSClaim{
			userEmail,
			jwt.StandardClaims{
				ExpiresAt: int64(claimExp),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaim)
		ret, err = token.SignedString(superSecret)
		if err != nil {
			util.LogError("error signing json web token: " + err.Error())
		} else {
			mySessions[ret] = true
		}
	}
	util.LogExit()
	return ret, err
}

// Validates the token. Returns {:email :exp} from token if session ok, None otherwise.
// Token validation has two parts:
// 1. Check that we actually created the token in the first place (should find it in my-sessions set.
// 2. Validate the actual token (can unsign it, token is not expired)."""
func ValidateJsonWebToken(myToken string) (ret TokenResponse, err error) {
	util.LogEnter()
	var parsedToken *jwt.Token
	var buf string
	// Validation #1.
	if _, ok := mySessions[myToken]; !ok {
		buf = "Token not found in sessions: " + myToken
		util.LogWarn(buf)
		err = errors.New(buf)
	} else {
		// Validation #2.
		parsedToken, err = jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
			return superSecret, nil
		})
		if err != nil {
			util.LogError("Couldn't parse token, error: " + err.Error())
		} else {
			claim, ok := parsedToken.Claims.(jwt.MapClaims) // ; ok && token.Valid
			if !ok {
				buf = "Couldn't parse token, Claims returned false"
				util.LogError(buf)
				err = errors.New(buf)
			} else {
				if !parsedToken.Valid {
					buf = "Token was not valid, parsedToken.Valid is false"
					util.LogError(buf)
					err = errors.New(buf)
				} else {
					userEmail := claim["email"]
					userEmailStr, ok := userEmail.(string)
					if !ok {
						buf = "Couldn't convert userEmail to string"
						util.LogError(buf)
						err = errors.New(buf)
					} else {
						ret = TokenResponse{true, userEmailStr}
					}
				}
			}
		}
	}
	util.LogExit()
	return ret, err
}
