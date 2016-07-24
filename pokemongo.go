package pokemongo

import (
	"bytes"
	"errors"
	"fmt"
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
	//apiURL = "https://httpbin.org/post"
)

type PokemonGo struct {
	client       *http.Client
	connector    auth.PokemonGoConnector
	playerCoords s2.LatLng
	authTicket   envelope.AuthTicket
	apiEndPoint  string
}

func (pg *PokemonGo) Execute(requestMethod enum.RequestMethod, request proto.Message) ([]byte, error) {
	reqEnv, err := pg.prepareRequest(requestMethod, request)
	if err != nil {
		return []byte{}, err
	}
	respEnv, err := pg.sendRequest(reqEnv)
	if err != nil {
		return []byte{}, err
	}
	pg.updateState(respEnv)
	err = chechResponse(respEnv)
	if err == errIncorrectAPIEndpoint {
		return pg.Execute(requestMethod, request)
	}
	if err != nil {
		return []byte{}, err
	}
	return respEnv.Responses[0], nil
}

func (pg *PokemonGo) prepareRequest(requestMethod enum.RequestMethod, request proto.Message) (envelope.Request, error) {
	raw, err := proto.Marshal(request)
	if err != nil {
		return envelope.Request{}, err
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
	unknown12 := unknown12
	reqEnv := envelope.Request{
		Direction: &direction,
		RpcId:     &requestID,
		Requests:  reqs,
		Latitude:  &latitude,
		Longitude: &longitude,
		Altitude:  &altitude,
		Unknown12: &unknown12,
	}
	if authTicketValid(pg.authTicket) {
		reqEnv.AuthTicket = &pg.authTicket
	} else {
		info, err := pg.connector.AuthInfo()
		if err != nil {
			return envelope.Request{}, err
		}
		authInfo := adaptAuthInfo(info)
		reqEnv.Auth = &authInfo
	}
	return reqEnv, nil
}

func (pg *PokemonGo) sendRequest(request envelope.Request) (envelope.Response, error) {
	raw, err := proto.Marshal(&request)
	if err != nil {
		return envelope.Response{}, err
	}
	httpReq, err := http.NewRequest("POST", pg.apiEndPoint, bytes.NewReader(raw))
	if err != nil {
		return envelope.Response{}, err
	}
	httpclient.AddUserAgent(httpReq)
	resp, err := pg.client.Do(httpReq)
	defer resp.Body.Close()
	if err != nil {
		return envelope.Response{}, err
	}
	if resp.StatusCode != http.StatusOK {
		return envelope.Response{}, fmt.Errorf("Got status %v for request", resp.StatusCode)
	}
	raw, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return envelope.Response{}, err
	}
	var respEnv envelope.Response
	err = proto.Unmarshal(raw, &respEnv)
	if err != nil {
		return envelope.Response{}, err
	}
	return respEnv, nil
}

func authTicketValid(authTicket envelope.AuthTicket) bool {
	return authTicket.ExpireTimestampMs != nil && *authTicket.ExpireTimestampMs > uint64(util.TimeInMilliseconds())
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

func (pg *PokemonGo) updateState(response envelope.Response) {
	if response.ApiUrl != nil && *response.ApiUrl != "" {
		endpoint := "https://" + *response.ApiUrl + "/rpc"
		pg.apiEndPoint = endpoint
	}
	if response.AuthTicket != nil {
		pg.authTicket = *response.AuthTicket
	}
}

var errIncorrectAPIEndpoint = errors.New("Incorrect API endpoint")

func chechResponse(response envelope.Response) error {
	if *response.Direction == 102 {
		return errors.New("Login failed")
	}
	if *response.Direction == 53 {
		return errIncorrectAPIEndpoint
	}
	if len(response.Responses) != 1 {
		return errors.New("Incorrect number of response")
	}
	return nil
}

func NewPokemonGo(connector auth.PokemonGoConnector, client *http.Client, playerCoords s2.LatLng) PokemonGo {
	return PokemonGo{
		client:       client,
		connector:    connector,
		playerCoords: playerCoords,
		apiEndPoint:  apiURL,
	}
}
