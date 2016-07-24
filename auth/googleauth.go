package auth

import (
	"net/http"

	"github.com/AKosterin/goandroidauth"
)

const (
	googleLoginAndroiAPI = "9774d56d682e549c"
	googleLoginService   = "audience:server:client_id:848232511240-7so421jotr2609rmqakceuu1luuq0ptb.apps.googleusercontent.com"
	googleLoginApp       = "com.nianticlabs.pokemongo"
	googleLoginClientSig = "321187995bc7cdc2b5fc91b11a96e2baa8602c62"
)

type googleConnector struct {
	info AuthInfo
}

func (c *googleConnector) AuthInfo() (AuthInfo, error) {
	return c.info, nil
}

func newGoogleConnector(ld LoginDetails, client *http.Client) (googleConnector, error) {
	tkn, err := getPTCToken(ld, client)
	googleConnect := googleConnector{AuthInfo{Provider: GOOGLE, Token: tkn}}
	return googleConnect, err
}

func getGoogleToken(ld LoginDetails) (Token, error) {
	naa := goandroidauth.NewAndroidAuth(googleLoginAndroiAPI, googleLoginApp, googleLoginClientSig, googleLoginService)

	token, err := naa.Login(ld.Username, ld.Password)

	return Token(token), err
}
