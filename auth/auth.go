package auth

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type Provider string

const (
	PTC    = Provider("ptc")
	GOOGLE = Provider("google")
)

type LoginDetails struct {
	Provider `json:"provider"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token string

type AuthInfo struct {
	Provider Provider
	Token    Token
}

type PokemonGoConnector interface {
	AuthInfo() (AuthInfo, error)
}

func LoginDetailsFromJSON(reader io.Reader) (LoginDetails, error) {
	raw, err := ioutil.ReadAll(reader)
	if err != nil {
		return LoginDetails{}, err
	}
	var loginDetails LoginDetails
	err = json.Unmarshal(raw, &loginDetails)
	return loginDetails, err
}
