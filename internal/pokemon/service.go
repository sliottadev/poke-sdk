package pokemon

import (
	"fmt"

	"github.com/sliottadev/poke-sdk/internal/httpclient"
)

func GetPokemon(client *httpclient.HTTPClient, baseURL string, nameOrID string) (*Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", baseURL, nameOrID)

	var pokemon Pokemon
	if err := client.Get(url, &pokemon); err != nil {
		return nil, err
	}

	return &pokemon, nil
}
