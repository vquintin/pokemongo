package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/vquintin/pokemongo/httpclient"
	"github.com/AKosterin/goandroidauth"
)

const (
	GOOGLE_LOGIN_ANROID_ID = "9774d56d682e549c"
	GOOGLE_LOGIN_SERVICE = "audience:server:client_id:848232511240-7so421jotr2609rmqakceuu1luuq0ptb.apps.googleusercontent.com"
	GOOGLE_LOGIN_APP = "com.nianticlabs.pokemongo"
	GOOGLE_LOGIN_CLIENT_SIG = "321187995bc7cdc2b5fc91b11a96e2baa8602c62"
)

type GoogleLoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type GoogleConnector struct {
	info AuthInfo
}

func (c *GoogleConnector) AuthInfo() (AuthInfo, error) {
	return c.info, nil
}

func NewGoogleConnector(ld GoogleLoginDetails, client *http.Client) (PTCConnector, error) {
	tkn, err := getPTCToken(ld, client)
	ptcConnect := GoogleConnector{AuthInfo{Provider: "google", Token: tkn}}
	return ptcConnect, err
}

func getGoogleToken(ld GoogleLoginDetails) (Token, error) {
	naa := goandroidauth.NewAndroidAuth(GOOGLE_LOGIN_ANROID_ID, GOOGLE_LOGIN_APP, GOOGLE_LOGIN_CLIENT_SIG, GOOGLE_LOGIN_SERVICE)

	token, err := naa.Login(ld.Username, ld.Password)

	return Token(token), err
}
