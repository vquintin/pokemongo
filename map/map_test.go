package pokemongo

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/httpclient"
	"github.com/vquintin/pokemongo/rpc"
)

func TestGetCatchablePokemons(t *testing.T) {
	raw, err := ioutil.ReadFile("ptcLoginDetails.json")
	assert.NoError(t, err, "An error occured while reading the ptc login details")
	var loginDetails auth.PTCLoginDetails
	err = json.Unmarshal(raw, &loginDetails)
	assert.NoError(t, err, "An error occured while decoding the ptc login details")
	client := httpclient.NewClient()
	connector, err := auth.NewPTCConnector(loginDetails, client)
	assert.NoError(t, err, "Could not connect")
	authInfo, err := connector.AuthInfo()
	assert.NoError(t, err, "Error occured while retrieving log info")
	t.Logf("Auth info: %v", authInfo)

	lat := 48.8462
	lng := 2.3372
	latLng := s2.LatLngFromDegrees(lat, lng)
	pg := rpc.NewPokemonGo(&connector, client)

	m := NewPokemonMap(pg)
	poks, err := m.CatchablePokemons(latLng, 3)
	assert.NoError(t, err, "An error occured while retrieving the catchable pokemons")
	t.Logf("Pokemons:\n")
	for _, v := range poks {
		t.Logf("%v\n", v)
	}
}
