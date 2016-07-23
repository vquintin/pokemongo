package pokemongo

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/httpclient"
	"github.com/vquintin/pokemongo/protobuf/networking/requests"
	"github.com/vquintin/pokemongo/protobuf/networking/requests/messages"
)

func TestCanCommunicateWithPokemonGoAPI(t *testing.T) {
	raw, err := ioutil.ReadFile("ptcLoginDetails.json")
	assert.NoError(t, err, "An error occured while reading the ptc login details")
	var loginDetails auth.PTCLoginDetails
	err = json.Unmarshal(raw, &loginDetails)
	assert.NoError(t, err, "An error occured while decoding the ptc login details")
	client := httpclient.NewClient()
	connector, err := auth.NewPTCConnector(loginDetails, client)
	assert.NoError(t, err, "Could not connect")

	lat := 48.846944
	lng := 2.336944
	pg := NewPokemonGo(&connector, client, s2.LatLngFromDegrees(lat, lng))

	raw, err = pg.Execute(requests.RequestType_GET_PLAYER, &messages.GetPlayerMessage{})

}
