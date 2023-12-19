package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c *Client) FetchPokemon(pokemonName string) (PokemonResponse, error) {
	endpoint := "/pokemon"
	fullURL := baseURL + endpoint + fmt.Sprintf("/%s", pokemonName)

	// req, err := http.NewRequest("GET", fullURL, nil)
	// if err != nil {
	// 	return PokemonResponse{}, err
	// }

	// resp, err := c.httpClient.Do(req)
	// if err != nil {
	// 	return PokemonResponse{}, err
	// }

	dat, ok := c.cache.Get(fullURL)
	if ok {
		pokemonResponse := PokemonResponse{}
		err := json.Unmarshal(dat, &pokemonResponse)
		if err != nil {
			return PokemonResponse{}, err
		}

		return pokemonResponse, nil
	}

	resp, err := c.httpClient.Get(fullURL)
	if err != nil {
		return PokemonResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return PokemonResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	dat, err = io.ReadAll(resp.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemonResponse := PokemonResponse{}
	err = json.Unmarshal(dat, &pokemonResponse)
	if err != nil {
		return PokemonResponse{}, err
	}

	c.cache.Add(fullURL, dat)

	return pokemonResponse, nil
}
