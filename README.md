# Pokedex CLI

A command-line Pokedex application written in Go that allows you to explore the Pokemon world, catch Pokemon, and manage your collection using the PokeAPI.

## Features

- **Explore Locations:** Navigate through different regions and areas.
- **Catch Pokemon:** Try your luck at catching Pokemon found in the wild.
- **Pokedex Management:** View details of your caught Pokemon and list your entire collection.
- **Caching:** Efficiently caches API responses to reduce network requests and improve performance.

## Installation

Ensure you have [Go](https://golang.org/doc/install) installed on your system.

1. Clone the repository:
   ```bash
   git clone https://github.com/mandarvu/pokedex_go.git
   cd pokedex_go
   ```

2. Build the application:
   ```bash
   go build -o pokedex
   ```

3. Run the application:
   ```bash
   ./pokedex
   ```

## Usage

Once the application is running, you can use the following commands:

- `help`: Displays a help message with all available commands.
- `map`: Displays the names of 20 location areas in the Pokemon world. Each subsequent call displays the next 20 areas.
- `mapb`: Displays the previous 20 location areas.
- `explore <area_name>`: Lists all the Pokemon found in a specific location area.
- `catch <pokemon_name>`: Attempts to catch a Pokemon. Success depends on the Pokemon's base experience.
- `inspect <pokemon_name>`: Displays details (stats, types, height, weight) of a Pokemon you have caught.
- `pokedex`: Lists all the Pokemon you have successfully caught.
- `exit`: Exits the Pokedex.

## Project Structure

- `main.go`: Entry point of the application.
- `internal/commands/`: Implementation of CLI commands.
- `internal/pokecache/`: Caching logic for API requests.
- `internal/json/`: Struct definitions for PokeAPI JSON responses.
- `internal/input/`: Input parsing and cleaning utilities.

## License

[MIT](LICENSE)
