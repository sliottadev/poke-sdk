package main

import (
	"fmt"
	"log"
	"time"

	pokeapi "github.com/sliottadev/poke-sdk"
)

func main() {
	config := pokeapi.NewConfig().
		WithRetries(3).
		WithRetryInterval(500 * time.Millisecond).
		WithTimeout(1 * time.Second)

	client := pokeapi.NewClient(config)

	pokemon, err := client.GetPokemon("778")
	if err != nil {
		log.Fatalf("error obteniendo pokemon: %v", err)
	}

	fmt.Printf("Pokemon: ID=%d, Name=%s, Height=%d, Weight=%d\n",
		pokemon.ID, pokemon.Name, pokemon.Height, pokemon.Weight)
}
