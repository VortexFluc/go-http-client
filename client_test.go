package client

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestClientCanHitAPI(t *testing.T) {
	t.Run("Happy path - can hit the api and return a pokemon", func (*testing.T) {
		myClient := NewClient()
		poke, err := myClient.GetPokemonByName(context.Background(), "pikachu")
		assert.NoError(t, err)
		assert.Equal(t, "pikachu", poke.Name)
	})

	t.Run("Sad path - return an error when the pokemon does not exist", func (*testing.T) {
		myClient := NewClient()
		_, err := myClient.GetPokemonByName(context.Background(), "non-existant-pokemon")
		assert.Error(t, err)
	})

	t.Run("Happy path - testing the WithApiURL option function", func(*testing.T) {
		myClient := NewClient(
			WithApiURL("my-test-url"),
		)
		assert.Equal(t, "my-test-url", myClient.apiURL)
	})

	t.Run("Happy path - tests with httpClient works", func(*testing.T) {
		myClient := NewClient(
			WithApiURL("my-test-url"),
			WithHTTPClient(&http.Client{
				Timeout: 1 * time.Second,
			}),
		)

		assert.Equal(t, "my-test-url", myClient.apiURL)
		assert.Equal(t, 1 * time.Second, myClient.httpClient.Timeout)
	})
}