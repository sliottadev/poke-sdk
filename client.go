package pokeapi

import (
	"github.com/sliottadev/poke-sdk/internal/httpclient"
	"github.com/sliottadev/poke-sdk/internal/pokemon"
)

// PokeClient defines the interface for interacting with the PokeAPI service.
type PokeClient interface {
	// GetPokemon retrieves a Pokemon by name or ID.
	GetPokemon(nameOrID string) (*pokemon.Pokemon, error)
}

// Client is the default HTTP implementation of the PokeClient interface.
type Client struct {
	baseURL    string
	httpClient *httpclient.HTTPClient
}

var _ PokeClient = (*Client)(nil)

// NewClient creates a new Client with the given configuration.
func NewClient(cfg *Config) *Client {
	return &Client{
		baseURL:    "https://pokeapi.co/api/v2",
		httpClient: httpclient.NewHTTPClient(cfg.timeout, cfg.retries, cfg.RetryInterval),
	}
}

// GetPokemon retrieves a Pokemon by its name or ID.
func (c *Client) GetPokemon(nameOrID string) (*pokemon.Pokemon, error) {
	return pokemon.GetPokemon(c.httpClient, c.baseURL, nameOrID)
}
