package pokemongo

import "testing"
import "github.com/stretchr/testify/assert"

func Test1stPokemonIsBulbasaur(t *testing.T) {
	bulbasaur, err := NewPokemon(PokeID(1))
	assert.NoError(t, err, "Can't create Bulbasaur")
	assert.Equal(t, "Bulbasaur", bulbasaur.Name(), "Bulbasaur is not called Bulbasaur")
	t.Log(bulbasaur.String())
}

func Test151thPokemonIsMew(t *testing.T) {
	mew, err := NewPokemon(PokeID(151))
	assert.NoError(t, err, "Can't create Mew")
	assert.Equal(t, "Mew", mew.Name(), "Mew is not called Mew")
	t.Log(mew.String())
}
