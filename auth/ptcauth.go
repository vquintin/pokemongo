package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/vquintin/pokemongo"
)

const (
	loginURL        = "https://sso.pokemon.com/sso/login?service=https://sso.pokemon.com/sso/oauth2.0/callbackAuthorize"
	loginOAuth      = "https://sso.pokemon.com/sso/oauth2.0/accessToken"
	ticketRe        = ".*ticket="
	tokenRe1        = "&expires.*"
	tokenRe2        = ".*access_token="
	ptcClientSecret = "w8ScCUXJQc6kXKw8FiOhd8Fixzht18Dq3PEVkUCP5ZPxtgyWsbTvWHFLm2wNY0JR"
)

type PTCLoginDetails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PTCConnector struct {
	info AuthInfo
}

func (c *PTCConnector) AuthInfo() (AuthInfo, error) {
	return c.info, nil
}

func NewPTCConnector(ld PTCLoginDetails, client *http.Client) (PTCConnector, error) {
	tkn, err := getPTCToken(ld, client)
	ptcConnect := PTCConnector{AuthInfo{Provider: "ptc", Token: tkn}}
	return ptcConnect, err
}

type token string

type jData struct {
	Lt        string `json:"lt"`
	Execution string `json:"execution"`
}

type ticket string

func getPTCToken(ld PTCLoginDetails, client *http.Client) (Token, error) {
	jd, err := getJData(client)
	if err != nil {
		return "", err
	}
	tkt, err := getTicket(jd, ld, client)
	if err != nil {
		return "", err
	}
	tkn, err := getToken(tkt, client)
	return tkn, err
}

func getJData(client *http.Client) (jData, error) {
	req, err := http.NewRequest("GET", loginURL, nil)
	if err != nil {
		return jData{}, err
	}
	addUserAgent(req)
	resp, err := client.Do(req)
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

func getTicket(jd jData, ld PTCLoginDetails, client *http.Client) (ticket, error) {
	values := url.Values{
		"lt":        {jd.Lt},
		"execution": {jd.Execution},
		"_eventId":  {"submit"},
		"username":  {ld.Username},
		"password":  {ld.Password},
	}
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	addUserAgent(req)
	addContentTypeForm(req)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil && err.(*url.Error).Err != pokemongo.ErrSentinel {
		return "", err
	}
	re, err := regexp.Compile(ticketRe)
	if err != nil {
		return "", err
	}
	location := resp.Header.Get("Location")
	result := re.ReplaceAll([]byte(location), []byte{})
	return ticket(result), nil
}

func getToken(tkt ticket, client *http.Client) (Token, error) {
	values := url.Values{
		"client_id":     {"mobile-app_pokemon-go"},
		"redirect_uri":  {"https://www.nianticlabs.com/pokemongo/error"},
		"client_secret": {string(ptcClientSecret)},
		"grant_type":    {"refresh_token"},
		"code":          {string(tkt)},
	}
	req, err := http.NewRequest("POST", loginOAuth, strings.NewReader(values.Encode()))
	if err != nil {
		return "", err
	}
	addUserAgent(req)
	addContentTypeForm(req)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	re, err := regexp.Compile(tokenRe1)
	if err != nil {
		return "", err
	}
	result := re.ReplaceAll([]byte(raw), []byte{})
	re, err = regexp.Compile(tokenRe2)
	if err != nil {
		return "", err
	}
	result = re.ReplaceAll([]byte(result), []byte{})
	return Token(result), nil
}

func addUserAgent(req *http.Request) {
	req.Header.Add("User-Agent", "Niantic App")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
}

func addContentTypeForm(req *http.Request) {
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
}
