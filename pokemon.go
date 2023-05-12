package client

import "context"

func (c *Client) GetPokemonByName(ctx context.Context, pokemonName string) (Pokemon, error) {
	return Pokemon{}, nil
}