package pokemongo

import "errors"

type Pokemon struct {
	id uint
}

var ErrInvalidPokemon = errors.New("Invalid pokémon id")

func NewPokemon(id uint) (Pokemon, error) {
	if id == 0 || id > 151 {
		return Pokemon{}, ErrInvalidPokemon
	}
	return Pokemon{id}, nil
}
