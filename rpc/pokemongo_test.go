package rpc

import (
	"os"
	"testing"

	"github.com/golang/geo/s2"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo/auth"
	"github.com/vquintin/pokemongo/httpclient"
	"github.com/vquintin/pokemongo/protobuf/enum"
	"github.com/vquintin/pokemongo/protobuf/sub"
)

func TestCanCommunicateWithPokemonGoAPI(t *testing.T) {
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

	pg := NewPokemonGo(connector, client)
	playerCoords := s2.LatLngFromDegrees(48.846944, 2.336944)
	raw, err := pg.Execute(playerCoords, enum.RequestMethod_GET_PLAYER, &sub.GetPlayerRequest{})
	assert.NoError(t, err, "An error occured while executing the request")

	var playerResponse sub.GetPlayerResponse
	err = proto.Unmarshal(raw, &playerResponse)
	assert.NoError(t, err, "Could not get the player profile")
	assert.NotNil(t, playerResponse.Profile, "Player profile is empty")
	assert.NotNil(t, playerResponse.Profile.Username, "Player username is empty")
	t.Logf("Username: %v\n", *playerResponse.Profile.Username)
}
