package auth

import (
	"errors"
	"net/http"
)

var ErrUnknownProvider = errors.New("Unknown provider")

func NewConnector(ld LoginDetails, client *http.Client) (PokemonGoConnector, error) {
	if ld.Provider == GOOGLE {
		gc, err := newGoogleConnector(ld, client)
		return &gc, err
	}
	if ld.Provider == PTC {
		ptcc, err := newPTCConnector(ld, client)
		return &ptcc, err
	}
	return nil, ErrUnknownProvider
}
