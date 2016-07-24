package pokemongo

import (
	"fmt"

	"github.com/golang/geo/s2"
	"github.com/vquintin/pokemongo/common"
)

type CatchablePokemon struct {
	common.Pokemon
	s2.LatLng
}

type NearbyPokemon struct {
	common.Pokemon
	distance float64
}

func (np NearbyPokemon) Distance() float64 {
	return np.distance
}

func (np NearbyPokemon) String() string {
	return fmt.Sprintf("%v at %v meters", np.Pokemon, np.distance)
}
