# Poke-SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/sliottadev/poke-sdk.svg)](https://pkg.go.dev/github.com/sliottadev/poke-sdk)
[![Go Version](https://img.shields.io/badge/Go-1.21.4-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A lightweight and efficient Go SDK for interacting with the [PokeAPI](https://pokeapi.co/). This library provides a simple client to fetch Pokémon data, offering configurable options for timeouts and retries to ensure robust API calls.

## Installation

To install the `poke-sdk` library, use `go get`:

```bash
go get github.com/sliottadev/poke-sdk
```

## Usage

Here's a basic example demonstrating how to use the `poke-sdk` to fetch Pokémon data by its ID or name:

```go
package main

import (
	"fmt"
	"log"
	"time"

	pokeapi "github.com/sliottadev/poke-sdk"
)

func main() {
	// Configure the client with custom settings
	config := pokeapi.NewConfig().
		WithRetries(3).
		WithRetryInterval(500 * time.Millisecond).
		WithTimeout(1 * time.Second)

	client := pokeapi.NewClient(config)

	// Fetch a Pokémon by its ID or name
	pokemon, err := client.GetPokemon("778") // Or use "pikachu"
	if err != nil {
		log.Fatalf("Error fetching pokemon: %v", err)
	}

	fmt.Printf("Pokemon: ID=%d, Name=%s, Height=%d, Weight=%d\n",
		pokemon.ID, pokemon.Name, pokemon.Height, pokemon.Weight)
}
```

To run this example:

1.  Save the code above as `main.go` inside an `example/` directory at the root of your `poke-sdk` project.
2.  Navigate to the `example/` directory in your terminal.
3.  Run the example:
    ```bash
    go run main.go
    ```

## Configuration

The `poke-sdk` client can be configured using the `Config` struct and its builder methods:

*   **`NewConfig()`**: Creates a new `Config` instance with default values (10 seconds timeout, 3 retries, 500 milliseconds retry interval).
*   **`WithTimeout(duration time.Duration)`**: Sets the timeout for HTTP requests.
*   **`WithRetries(count int)`**: Sets the number of retry attempts for failed requests.
*   **`WithRetryInterval(duration time.Duration)`**: Sets the delay between retry attempts.

Example of custom configuration:

```go
config := pokeapi.NewConfig().
	WithRetries(5). // Try up to 5 times
	WithRetryInterval(2 * time.Second). // Wait 2 seconds between retries
	WithTimeout(30 * time.Second) // 30 second total timeout for each request
```

## Contributing

Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 