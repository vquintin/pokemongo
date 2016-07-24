package auth

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vquintin/pokemongo/httpclient"
)

func TestConnection(t *testing.T) {
	reader, err := os.Open("loginDetails.json")
	loginDetails, err := LoginDetailsFromJSON(reader)
	assert.NoError(t, err, "An error occured while decoding the login details")

	connector, err := NewConnector(loginDetails, httpclient.NewClient())

	assert.NoError(t, err, "An error occured while connecting")
	info, err := connector.AuthInfo()
	assert.NotEqual(t, "", info.Token, "The token is empty")
	t.Logf("Token: '%v'", info.Token)
}
