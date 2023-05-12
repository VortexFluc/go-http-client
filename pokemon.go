package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetPokemonByName(
	ctx context.Context,
	 pokemonName string,
) (Pokemon, error) {

	// Create a HTTP Request
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://pokeapi.co/api/v2/pokemon/"+pokemonName,
		nil,
	)
	// If creation of http request is not OK - return empty Pokemon and error
	if err != nil {
		return Pokemon{}, err
	}

	// Add header to request
	req.Header.Add("Accept", "application/json")

	// Proceed request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	// If not 200 - then return empty pokemon and error
	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("Unexpected status code returned from the pokeapi")
	}

	// Decode response body
	var pokemon Pokemon
	err = json.NewDecoder(resp.Body).Decode(&pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}