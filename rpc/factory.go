package rpc

import (
	"net/http"

	"github.com/vquintin/pokemongo/auth"
)

func NewPokemonGo(connector auth.PokemonGoConnector, client *http.Client) PokemonGo {
	return PokemonGo{
		client:      client,
		connector:   connector,
		apiEndPoint: apiURL,
	}
}
