package pokeapi

import (
	"github.com/sliottadev/poke-sdk/internal/httpclient"
	"github.com/sliottadev/poke-sdk/internal/pokemon"
)

type PokeClient interface {
	GetPokemon(nameOrID string) (*pokemon.Pokemon, error)
}

type Client struct {
	baseURL    string
	httpClient *httpclient.HTTPClient
}

var _ PokeClient = (*Client)(nil)

func NewClient(cfg *Config) *Client {
	return &Client{
		baseURL:    "https://pokeapi.co/api/v2",
		httpClient: httpclient.NewHTTPClient(cfg.timeout, cfg.retries, cfg.RetryInterval),
	}
}

func (c *Client) GetPokemon(nameOrID string) (*pokemon.Pokemon, error) {
	return pokemon.GetPokemon(c.httpClient, c.baseURL, nameOrID)
}
