package pokeapi

type LocationAreaResp struct {
	Next      *string			`json:"next"` 
	Prev 			*string			`json:"previous"`
	Locations []struct {
		Name string					`json:"name"`
		URL  string 				`json:"url"`	
	}  										`json:"results"`
}

type LocationEncounterResp struct {
	Name string `json:"name"`
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

