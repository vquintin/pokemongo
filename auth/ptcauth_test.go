package auth

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo"
)

func TestConnection(t *testing.T) {
	raw, err := ioutil.ReadFile("ptcLoginDetails.json")
	assert.NoError(t, err, "An error occured while reading the ptc login details")
	var loginDetails PTCLoginDetails
	err = json.Unmarshal(raw, &loginDetails)
	assert.NoError(t, err, "An error occured while decoding the ptc login details")

	connector, err := NewPTCConnector(loginDetails, pokemongo.NewClient())

	assert.NoError(t, err, "An error occured while connecting")
	info, err := connector.AuthInfo()
	assert.NotEqual(t, "", info.Token, "The token is empty")
	t.Logf("Token: '%v'", info.Token)
}
