package pokemongo

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"math"

	"github.com/golang/geo/s2"
	"github.com/golang/protobuf/proto"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/httpclient"
	"github.com/vquintin/pokemongo/protobuf/enum"
	"github.com/vquintin/pokemongo/protobuf/envelope"
	"github.com/vquintin/pokemongo/util"
)

const (
	requestID int64 = 8145806132888207460
	unknown12 int64 = 989
	unknown13 int32 = 59
	apiURL          = "https://pgorelease.nianticlabs.com/plfe/rpc"
)

type PokemonGo struct {
	client       *http.Client
	connector    auth.PokemonGoConnector
	playerCoords s2.LatLng
	authTicket   *envelope.AuthTicket
	apiEndPoint  string
}

func (pg *PokemonGo) Execute(requestMethod enum.RequestMethod, request proto.Message) ([]byte, error) {
	raw, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	req := envelope.Request_Requests{
		Type:       &requestMethod,
		Parameters: raw,
	}
	reqs := make([]*envelope.Request_Requests, 1)
	reqs[0] = &req
	direction := enum.RpcDirection_REQUEST
	requestID := requestID
	latitude := math.Float64bits(pg.playerCoords.Lat.Degrees())
	longitude := math.Float64bits(pg.playerCoords.Lng.Degrees())
	altitude := uint64(0)
	reqEnv := envelope.Request{
		Direction: &direction,
		RpcId:     &requestID,
		Requests:  reqs,
		Latitude:  &latitude,
		Longitude: &longitude,
		Altitude:  &altitude,
	}
	if authTicketValid(pg.authTicket) {
		reqEnv.AuthTicket = pg.authTicket
	} else {
		info, err := pg.connector.AuthInfo()
		if err != nil {
			return nil, err
		}
		authInfo := adaptAuthInfo(info)
		reqEnv.Auth = &authInfo
	}
	raw, err = proto.Marshal(&reqEnv)
	if err != nil {
		return nil, err
	}
	httpReq, err := http.NewRequest("POST", pg.apiEndPoint, bytes.NewReader(raw))
	if err != nil {
		return nil, err
	}
	httpclient.AddUserAgent(httpReq)
	resp, err := pg.client.Do(httpReq)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	raw, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var respEnv envelope.Response
	err = proto.Unmarshal(raw, &respEnv)
	if err != nil {
		return nil, err
	}
	if respEnv.ApiUrl != nil && *respEnv.ApiUrl != "" {
		pg.apiEndPoint = *respEnv.ApiUrl
	}
	if respEnv.AuthTicket != nil {
		pg.authTicket = respEnv.AuthTicket
	}
	if *respEnv.Direction == 102 {
		return nil, errors.New("Login failed")
	}
	if *respEnv.Direction == 53 {
		return nil, errors.New("API endpoint not correctly set")
	}
	if len(respEnv.Responses) != 1 {
		return nil, errors.New("Incorrect number of response")
	}
	return respEnv.Responses[0], nil
}

func authTicketValid(authTicket *envelope.AuthTicket) bool {
	return authTicket != nil && *authTicket.ExpireTimestampMs > uint64(util.TimeInMilliseconds())
}

func adaptAuthInfo(authInfo auth.AuthInfo) envelope.Request_AuthInfo {
	provider := string(authInfo.Provider)
	token := string(authInfo.Token)
	unknown13 := unknown13
	return envelope.Request_AuthInfo{
		Provider: &provider,
		Token: &envelope.Request_AuthInfo_JWT{
			Contents:  &token,
			Unknown13: &unknown13,
		},
	}
}

func NewPokemonGo(connector auth.PokemonGoConnector, client *http.Client, playerCoords s2.LatLng) PokemonGo {
	return PokemonGo{
		client:       client,
		connector:    connector,
		playerCoords: playerCoords,
		apiEndPoint:  apiURL,
	}
}
