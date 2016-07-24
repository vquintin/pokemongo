package pokemongo

import (
	"fmt"
	"sort"
	"time"

	"github.com/golang/geo/s2"
	"github.com/golang/protobuf/proto"
	"github.com/vquintin/pokemongo/common"
	"github.com/vquintin/pokemongo/protobuf/enum"
	"github.com/vquintin/pokemongo/protobuf/sub"
	"github.com/vquintin/pokemongo/rpc"
	"github.com/vquintin/pokemongo/util"
)

type Map struct {
	rpc         rpc.PokemonGo
	lastUpdates map[s2.CellID]time.Time
}

func (pkmnMap *Map) CatchablePokemons(coords s2.LatLng, width uint) ([]CatchablePokemon, error) {
	mapResp, err := pkmnMap.fetchMapObjectsResponse(coords, width)
	if err != nil {
		return []CatchablePokemon{}, err
	}
	return catchablePokemons(mapResp)
}

func catchablePokemons(resp sub.GetMapObjectsResponse) ([]CatchablePokemon, error) {
	var catchables []CatchablePokemon
	for _, mapCell := range resp.GetMapCells() {
		for _, catchable := range mapCell.CatchablePokemons {
			pok, err := makePokemon(*catchable)
			if err != nil {
				return []CatchablePokemon{}, err
			}
			catchables = append(catchables, pok)
		}
	}
	return catchables, nil
}

func makePokemon(pok sub.MapPokemon) (CatchablePokemon, error) {
	latLng := s2.LatLngFromDegrees(*pok.Latitude, *pok.Longitude)
	pokemon, err := common.NewPokemon(common.PokeID(*pok.PokemonId))
	if err != nil {
		return CatchablePokemon{}, err
	}
	return CatchablePokemon{
		Pokemon: pokemon,
		LatLng:  latLng,
	}, nil
}

func (pkmnMap *Map) fetchMapObjectsResponse(coords s2.LatLng, width uint) (sub.GetMapObjectsResponse, error) {
	cells := makeCellIds(coords, width)
	cellIds := func(cells []s2.CellID) []uint64 {
		result := []uint64{}
		for _, v := range cells {
			result = append(result, uint64(v))
		}
		return result
	}(cells)
	latitude := coords.Lat.Degrees()
	longitude := coords.Lng.Degrees()
	message := sub.GetMapObjectsRequest{
		CellId:           cellIds,
		SinceTimestampMs: pkmnMap.getLastUpdatesInMs(cells),
		Latitude:         &latitude,
		Longitude:        &longitude,
	}
	fmt.Println(message)
	raw, err := pkmnMap.rpc.Execute(coords, enum.RequestMethod_GET_MAP_OBJECTS, &message)
	if err != nil {
		return sub.GetMapObjectsResponse{}, err
	}
	var mapResp sub.GetMapObjectsResponse
	err = proto.Unmarshal(raw, &mapResp)
	if err != nil {
		return sub.GetMapObjectsResponse{}, err
	}
	return mapResp, nil
}

type s2Cells []s2.CellID

func (c s2Cells) Len() int {
	return len(c)
}

func (c s2Cells) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c s2Cells) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func makeCellIds(latLng s2.LatLng, width uint) []s2.CellID {
	cellId := s2.CellIDFromLatLng(latLng).Parent(15)
	result := s2Cells{}
	result = append(result, cellId)
	sort.Sort(result)
	return result
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

func NewPokemonMap(rpc rpc.PokemonGo) Map {
	return Map{rpc: rpc}
}
