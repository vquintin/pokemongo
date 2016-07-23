package pokemongo

import (
	"time"

	"github.com/golang/geo/s2"
	"github.com/vquintin/pokemongo/protobuf/networking/requests/messages"
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
	result := make([]s2.CellID, 0, 9)
	result = append(result, cellId)
	result = append(result, cellId.VertexNeighbors(15)...)
	for _, v := range cellId.EdgeNeighbors() {
		result = append(result, v)
	}
	return result
}

func (pkmnMap *Map) fetchMapObjects(cells []s2.CellID) (MapObjects, error) {
	cellIds := func(cells []s2.CellID) []uint64 {
		result := make([]uint64, len(cells))
		for _, v := range cells {
			result = append(result, uint64(v))
		}
		return result
	}(cells)
	message := messages.GetMapObjectsMessage{
		CellId:           cellIds,
		SinceTimestampMs: pkmnMap.getLastUpdatesInMs(cells),
		Latitude:         pkmnMap.api.playerCoords.Lat.Degrees(),
		Longitude:        pkmnMap.api.playerCoords.Lng.Degrees(),
	}
	fmt.Prinfln(message)
	return MapObjects{}, err
}

func (pkmnMap Map) getLastUpdatesInMs(cells []s2.CellID) []int64 {
	result := make([]int64, len(cells))
	for _, v := range cells {
		date := pkmnMap.lastUpdates[v]
		millis := util.ConvertToMilliseconds(date)
		result = append(result, millis)
	}
	return result
}
