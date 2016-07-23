package pokemongo

import (
	"net/http"

	"github.com/golang/geo/s2"
	"github.com/golang/protobuf/proto"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/protobuf/networking/envelopes"
	"github.com/vquintin/pokemongo/protobuf/networking/requests"
)

const (
	statusCode int32  = 2
	requestID  uint64 = 8145806132888207460
	unknown12  int64  = 989
)

type PokemonGo struct {
	client       *http.Client
	connector    auth.PokemonGoConnector
	playerCoords s2.LatLng
	authTicket   *envelopes.AuthTicket
}

func (pg *PokemonGo) Execute(requestType requests.RequestType, request proto.Message) (interface{}, error) {
	raw, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	req := requests.Request{
		RequestType:    requests.RequestType_GET_MAP_OBJECTS,
		RequestMessage: raw,
	}
	reqs := make([]*requests.Request, 1)
	reqs[0] = &req
	reqEnv := envelopes.RequestEnvelope{
		StatusCode: statusCode,
		RequestId:  requestID,
		Requests:   reqs,
		Latitude:   pg.playerCoords.Lat.Degrees(),
		Longitude:  pg.playerCoords.Lng.Degrees(),
		Altitude:   0,
	}
}

func authTicketValid(authTicket *envelopes.AuthTicket) bool {
	return authTicket != nil && authTicket.
}

func timeInMillis() {
    
}

func NewPokemonGo(connector auth.PokemonGoConnector, client *http.Client, playerCoords s2.LatLng) {

}
