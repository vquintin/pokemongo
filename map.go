package pokemongo

import (
	"fmt"
	"time"

	"github.com/golang/geo/s2"
	"github.com/golang/protobuf/proto"
	"github.com/vquintin/pokemongo/protobuf/enum"
	"github.com/vquintin/pokemongo/protobuf/sub"
	"github.com/vquintin/pokemongo/util"
)

type CatchablePokemon struct {
	Pokemon
	s2.LatLng
}

type Map struct {
	api         PokemonGo
	lastUpdates map[s2.CellID]time.Time
}

func (pkmnMap *Map) CatchablePokemons() ([]CatchablePokemon, error) {
	return []CatchablePokemon{}, nil
}

func (pkmnMap *Map) mapObjects() (MapObjects, error) {
	return MapObjects{}, nil
}

func (pkmnMap *Map) makeCellIds(latLng s2.LatLng) []s2.CellID {
	cellId := s2.CellIDFromLatLng(latLng).Parent(15)
	result := []s2.CellID{}
	result = append(result, cellId)
	result = append(result, cellId.VertexNeighbors(15)...)
	for _, v := range cellId.EdgeNeighbors() {
		result = append(result, v)
	}
	return result
}

func (pkmnMap *Map) fetchMapObjects(cells []s2.CellID) (MapObjects, error) {
	cellIds := func(cells []s2.CellID) []uint64 {
		result := []uint64{}
		for _, v := range cells {
			result = append(result, uint64(v))
		}
		return result
	}(cells)
	latitude := pkmnMap.api.playerCoords.Lat.Degrees()
	longitude := pkmnMap.api.playerCoords.Lng.Degrees()
	message := sub.GetMapObjectsRequest{
		CellId:           cellIds,
		SinceTimestampMs: pkmnMap.getLastUpdatesInMs(cells),
		Latitude:         &latitude,
		Longitude:        &longitude,
	}
	fmt.Println(message)
	raw, err := pkmnMap.api.Execute(enum.RequestMethod_GET_MAP_OBJECTS, &message)
	if err != nil {
		return MapObjects{}, err
	}
	var mapResp sub.GetMapObjectsResponse
	err = proto.Unmarshal(raw, &mapResp)
	if err != nil {
		return MapObjects{}, err
	}
	var catchables []CatchablePokemon
	for _, mapCell := range mapResp.GetMapCells() {
		for _, catchable := range mapCell.CatchablePokemons {
			pok, err := makePokemon(*catchable)
			if err != nil {
				return MapObjects{}, err
			}
			catchables = append(catchables, pok)
		}
	}
	return MapObjects{catchables}, nil
}

func makePokemon(pok sub.MapPokemon) (CatchablePokemon, error) {
	latLng := s2.LatLngFromDegrees(*pok.Latitude, *pok.Longitude)
	pokemon, err := NewPokemon(PokeID(*pok.PokemonId))
	if err != nil {
		return CatchablePokemon{}, err
	}
	return CatchablePokemon{
		Pokemon: pokemon,
		LatLng:  latLng,
	}, nil
}

func (pkmnMap Map) getLastUpdatesInMs(cells []s2.CellID) []int64 {
	result := []int64{}
	for _, v := range cells {
		date := pkmnMap.lastUpdates[v]
		millis := util.ConvertToMilliseconds(date)
		result = append(result, millis)
	}
	return result
}

func NewPokemonMap(api PokemonGo) Map {
	return Map{api: api}
}
