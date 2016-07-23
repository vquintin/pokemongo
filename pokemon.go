package pokemongo

import "errors"

type Pokemon struct {
	id uint
}

var ErrInvalidPokemon = errors.New("Invalid pokémon id")

func NewPokemon(uint id) (Pokemon, err) {
	if id == 0 || id > 151 {
		return Pokemon{}, err
	}
	return Pokemon{id}
}
