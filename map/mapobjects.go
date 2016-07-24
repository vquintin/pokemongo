package pokemongo

import (
	"github.com/golang/geo/s2"
	"github.com/vquintin/pokemongo/common"
)

type CatchablePokemon struct {
	common.Pokemon
	s2.LatLng
}
