package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetPokemon -
func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	if val, ok := c.cache.Get(url); ok {
		pokemonResp := Pokemon{}
		err := json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error unmarshalling pokemon from cache: %v", err)
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error making GET request for Pokemon: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error making GET request of client for Pokemon: %v", err)
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error reading from Get request response for Pokemon: %v", err)
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshalling from Get request response for Pokemon: %v", err)
	}

	c.cache.Add(url, dat)

	return pokemonResp, nil
}
