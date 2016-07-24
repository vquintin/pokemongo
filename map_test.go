package pokemongo

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/httpclient"
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

	lat := 48.846944
	lng := 2.336944
	latLng := s2.LatLngFromDegrees(lat, lng)
	pg := NewPokemonGo(&connector, client, s2.LatLngFromDegrees(lat, lng))

	m := NewPokemonMap(pg)

	cell := s2.CellIDFromLatLng(latLng).Parent(15)

	mo, err := m.fetchMapObjects([]s2.CellID{cell})
	assert.NoError(t, err, "An error occured while retrieving the map objects")
	t.Logf("Pokemons:\n")
	for _, v := range mo.CatchablePokemons() {
		t.Logf("%v\n", v)
	}
}
