package pokemongo

import (
	"os"
	"testing"

	"github.com/golang/geo/s2"
	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/httpclient"
	"github.com/vquintin/pokemongo/rpc"
)

func TestGetCatchablePokemons(t *testing.T) {
	reader, err := os.Open("loginDetails.json")
	defer reader.Close()
	assert.NoError(t, err, "An error occured while reading the login details")
	loginDetails, err := auth.LoginDetailsFromJSON(reader)
	assert.NoError(t, err, "An error occured while decoding the login details")
	client := httpclient.NewClient()
	connector, err := auth.NewConnector(loginDetails, client)
	assert.NoError(t, err, "Could not connect")
	authInfo, err := connector.AuthInfo()
	assert.NoError(t, err, "Error occured while retrieving log info")
	t.Logf("Auth info: %v", authInfo)

	lat := 48.8462
	lng := 2.3372
	pg := rpc.NewPokemonGo(connector, client)

	m := NewPokemonMap(pg)
	latLng := s2.LatLngFromDegrees(lat, lng)
	poks, err := m.CatchablePokemons(latLng)
	assert.NoError(t, err, "An error occured while retrieving the catchable pokemons")
	t.Logf("Pokemons:\n")
	for _, v := range poks {
		t.Logf("%v\n", v)
	}
}

func TestGetNearbyPokemons(t *testing.T) {
	reader, err := os.Open("loginDetails.json")
	defer reader.Close()
	assert.NoError(t, err, "An error occured while reading the login details")
	loginDetails, err := auth.LoginDetailsFromJSON(reader)
	assert.NoError(t, err, "An error occured while decoding the login details")
	client := httpclient.NewClient()
	connector, err := auth.NewConnector(loginDetails, client)
	assert.NoError(t, err, "Could not connect")
	authInfo, err := connector.AuthInfo()
	assert.NoError(t, err, "Error occured while retrieving log info")
	t.Logf("Auth info: %v", authInfo)

	lat := 48.8462
	lng := 2.3372
	pg := rpc.NewPokemonGo(connector, client)

	m := NewPokemonMap(pg)
	latLng := s2.LatLngFromDegrees(lat, lng)
	poks, err := m.NearbyPokemons(latLng, 10)
	assert.NoError(t, err, "An error occured while retrieving the catchable pokemons")
	t.Logf("Pokemons:\n")
	for _, v := range poks {
		t.Logf("%v\n", v)
	}
}
