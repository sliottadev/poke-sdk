package pokeapi

import "github.com/sliottadev/poke-sdk/internal/pokemon"

type ClientMock struct {
	getPokemonFunc func(nameOrID string) (*pokemon.Pokemon, error)
}

var _ PokeClient = (*ClientMock)(nil)

func NewClientMock() *ClientMock {
	return &ClientMock{}
}

func (m *ClientMock) WithGetPokemon(f func(nameOrID string) (*pokemon.Pokemon, error)) *ClientMock {
	m.getPokemonFunc = f
	return m
}

func (m *ClientMock) GetPokemon(nameOrID string) (*pokemon.Pokemon, error) {
	if m.getPokemonFunc != nil {
		return m.getPokemonFunc(nameOrID)
	}
	return nil, nil
}
