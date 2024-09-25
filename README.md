# Pokedex CLI

A command-line interface Pokedex application built with Go, utilizing the PokeAPI.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Commands](#commands)
- [Project Structure](#project-structure)
- [Cache System](#cache-system)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This Pokedex CLI is a Go-based application that allows users to explore the world of Pokemon through a command-line interface. It interacts with the PokeAPI to fetch data about Pokemon, locations, and more.

## Features

- Explore Pokemon locations
- Catch Pokemon
- View caught Pokemon in your Pokedex
- Inspect Pokemon details
- Efficient caching system for API responses

## Installation

1. Ensure you have Go installed on your system (version 1.23.1 or later).
2. Clone this repository:

`git clone https://github.com/Jschles1/pokedexcli.git`
   
3. Navigate to the project directory:
   
`cd pokedexcli`
   
4. Build the application:

`go build`


## Usage

To start the Pokedex CLI, run the compiled binary:

`./pokedexcli`

Once started, you'll see a prompt where you can enter commands:

`Pokedex >`

## Commands

The following commands are available in the Pokedex CLI:

- `help`: Displays a help message with all available commands.
- `exit`: Exits the Pokedex CLI.
- `map`: Displays the names of 20 location areas in the Pokemon world.
- `mapb`: Displays the previous 20 location areas.
- `explore <area_name>`: Lists all the Pokemon in a given area.
- `catch <pokemon_name>`: Attempts to catch the specified Pokemon.
- `inspect <pokemon_name>`: Displays details of a caught Pokemon.
- `pokedex`: Lists all caught Pokemon.

For more details on each command, refer to the `commands.go` file.

## Project Structure

The project is structured as follows:

- `main.go`: Entry point of the application.
- `commands.go`: Defines and implements CLI commands.
- `internal/`:
  - `pokeapi/`: Contains the PokeAPI client for making API requests.
  - `pokecache/`: Implements the caching system.

## Cache System

The application uses an in-memory cache system to store API responses, improving performance and reducing API calls. The cache is implemented in the `pokecache` package.

## Testing

The project includes unit tests for the cache system. To run the tests, use the following command:

`go test ./...`

## Contributing

Contributions to the Pokedex CLI are welcome! Please feel free to submit issues, fork the repository and send pull requests.

1. Fork the project
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a pull request

## License

This project is open source and available under the [MIT License](LICENSE).