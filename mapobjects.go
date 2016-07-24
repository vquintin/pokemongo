package pokemongo

type MapObjects struct {
	catchables []CatchablePokemon
}

func (mo MapObjects) CatchablePokemons() []CatchablePokemon {
	return mo.catchables
}
