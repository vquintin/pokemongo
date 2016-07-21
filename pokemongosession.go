package pokemongogo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

const (
	API_URL     = "https://pgorelease.nianticlabs.com/plfe/rpc"
	LOGIN_URL   = "https://sso.pokemon.com/sso/login?service=https://sso.pokemon.com/sso/oauth2.0/callbackAuthorize"
	LOGIN_OAUTH = "https://sso.pokemon.com/sso/oauth2.0/accessToken"
	APP         = "com.nianticlabs.pokemongo"
	TICKET_RE   = ""
	TOKEN_RE1   = ""
	TOKEN_RE2   = ""
)

type PTCClientSecret string
type AndroidId string
type Service string
type ClientSig string

type Credentials struct {
	PTCClientSecret `json:"ptcClientSecret"`
	AndroidId       `json:"androidId"`
	Service         `json:"service"`
	ClientSig       `json:"clienSig"`
}

type PokemonGoSession struct {
	credentials Credentials
	username    string
	password    string
	client      http.Client
	token
}

func NewPokemonGoSession(creds Credentials, username string, password string) (PokemonGoSession, error) {
	pgs := PokemonGoSession{credentials: creds, username: username, password: password}
	tkn, err := pgs.getPTCToken()
	pgs.token = tkn
	return pgs, err
}

type token string

type jData struct {
	execution string
	lt        string
}

type ticket string

func (pgs *PokemonGoSession) getPTCToken() (token, error) {
	jd, err := pgs.getJData()
	if err != nil {
		return "", err
	}
	tkt, err := pgs.getTicket(jd)
	if err != nil {
		return "", err
	}
	tkn, err := pgs.getToken(tkt)
	return tkn, err
}

func (pgs *PokemonGoSession) getJData() (jData, error) {
	req, err := http.NewRequest("GET", LOGIN_URL, nil)
	if err != nil {
		return jData{}, err
	}
	addUserAgent(req)
	resp, err := pgs.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return jData{}, err
	}
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return jData{}, err
	}
	var jd jData
	err = json.Unmarshal(raw, &jd)
	if err != nil {
		return jData{}, err
	}
	return jd, nil
}

func (pgs *PokemonGoSession) getTicket(jd jData) (ticket, error) {
	req, err := http.NewRequest("POST", LOGIN_URL, nil)
	if err != nil {
		return "", err
	}
	addUserAgent(req)
	req.PostForm.Add("lt", jd.lt)
	req.PostForm.Add("execution", jd.execution)
	req.PostForm.Add("_eventId", "submit")
	req.PostForm.Add("username", pgs.username)
	req.PostForm.Add("password", pgs.password)
	resp, err := pgs.client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
	re, err := regexp.Compile(TICKET_RE)
	if err != nil {
		return "", err
	}
	location := resp.Header.Get("Location")
	result := re.ReplaceAll([]byte(location), []byte{})
	return ticket(result), nil
}

func (pgs *PokemonGoSession) getToken(tkt ticket) (token, error) {
	req, err := http.NewRequest("POST", LOGIN_URL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("client_id", "mobile-app_pokemon-go")
	req.Header.Add("redirect_uri", "https://www.nianticlabs.com/pokemongo/error")
	req.Header.Add("client_secret", string(pgs.credentials.PTCClientSecret))
	req.Header.Add("grant_type", "refresh_token")
	req.Header.Add("code", string(tkt))
	resp, err := pgs.client.Do(req)
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re, err := regexp.Compile(TOKEN_RE1)
	if err != nil {
		return "", err
	}
	result := re.ReplaceAll([]byte(raw), []byte{})
	re, err = regexp.Compile(TOKEN_RE2)
	if err != nil {
		return "", err
	}
	result = re.ReplaceAll([]byte(result), []byte{})
	return token(result), nil
}

func addUserAgent(req *http.Request) {
	req.Header.Add("User-Agent", "Niantic App")
}

func debug(format string, a ...interface{}) (int, error) {
	return fmt.Printf(format, a)
}
