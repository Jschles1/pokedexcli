package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(string, *PokeAPIConfig, *map[string]Pokemon) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message.",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex.",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the next 20 locations, and so on.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the names of 20 location areas in the Pokemon world. Each subsequent call to map should display the previous 20 locations, and so on.",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "displays a list of all the Pokémon in a given area. Use the \"map\" and \"mapb\" commands to display valid areas.",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch given Pokemon with a pokeball. If successful, pokemon is added to user's pokedex.",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "Prints the name, height, weight, stats and type(s) of the Pokemon, if caught.",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Print a list of all the names of the Pokemon you have caught.",
			callback:    commandPokedex,
		},
	}
}

// Commands //

func commandHelp(param string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	commands := getCommands()
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	fmt.Println()

	for _, command := range commands {
		fmt.Printf("%v: %v", command.name, command.description)
		fmt.Println()
	}

	return nil
}

func commandExit(param string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	fmt.Println("Exiting Pokedex...")
	os.Exit(0)
	return nil
}

func commandMap(param string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	url := config.Next
	if url == "" {
		fmt.Println("No more locations available, please use mapb to go back")
		return nil
	}
	locationsResponse, err := config.Client.GetLocationsResponse(url)
	if err != nil {
		fmt.Printf("Error retrieving locations: %v", err)
		fmt.Println()
		return nil
	}
	config.Next = locationsResponse.Next
	if locationsResponse.Previous != nil {
		config.Previous = *locationsResponse.Previous
	} else {
		config.Previous = ""
	}
	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(param string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	url := config.Previous
	if url == "" {
		fmt.Println("No more previous locations available, please use map retrieve locations")
		return nil
	}
	locationsResponse, err := config.Client.GetLocationsResponse(url)
	if err != nil {
		fmt.Printf("Error retrieving locations: %v", err)
		fmt.Println()
		return nil
	}
	config.Next = locationsResponse.Next
	if locationsResponse.Previous != nil {
		config.Previous = *locationsResponse.Previous
	} else {
		config.Previous = ""
	}
	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandExplore(param string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	fmt.Println("Exploring " + param + "...")
	exploreResponse, err := config.Client.GetExploreLocationAreaResponse(param)
	if err != nil {
		fmt.Printf("Error retrieving pokémon from location: %v", err)
		fmt.Println()
		return nil
	}
	fmt.Println("Found Pokemon:")
	for _, result := range exploreResponse.PokemonEncounters {
		fmt.Println(" - " + result.Pokemon.Name)
	}
	return nil
}

func commandCatch(pokemon string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	pokemonResponse, err := config.Client.GetPokemonDetailResponse(pokemon)
	if err != nil {
		fmt.Printf("Error: %v", err)
		fmt.Println()
		return nil
	}
	pokemonName := pokemonResponse.Name
	fmt.Println("Throwing a Pokeball at " + pokemonName + "...")
	time.Sleep(3 * time.Second)
	roll := rand.Intn(pokemonResponse.BaseExperience) >= pokemonResponse.BaseExperience/2
	if roll {
		fmt.Println(pokemonName + " was caught!")
		stats := []PokemonStat{}
		for _, stat := range pokemonResponse.Stats {
			formattedStat := PokemonStat{
				name:  stat.Stat.Name,
				value: stat.BaseStat,
			}
			stats = append(stats, formattedStat)
		}
		types := []string{}
		for _, pokemonType := range pokemonResponse.Types {
			types = append(types, pokemonType.Type.Name)
		}
		caughtPokemon := Pokemon{
			name:   pokemonName,
			height: pokemonResponse.Height,
			weight: pokemonResponse.Weight,
			stats:  stats,
			types:  types,
		}
		(*pokedex)[pokemonName] = caughtPokemon
		fmt.Println("You may now inspect it with the inspect command.")
	} else {
		fmt.Println(pokemonName + " escaped!")
	}
	return nil
}

func commandInspect(pokemon string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	p, ok := (*pokedex)[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	} else {
		fmt.Println("Name: ", p.name)
		fmt.Println("Height: ", strconv.Itoa(p.height))
		fmt.Println("Weight: ", strconv.Itoa(p.weight))
		fmt.Println("Stats:")
		for _, stat := range p.stats {
			fmt.Println("-" + stat.name + ": " + strconv.Itoa(stat.value))
		}
		fmt.Println("Types:")
		for _, pokemonType := range p.types {
			fmt.Println("- " + pokemonType)
		}
	}
	return nil
}

func commandPokedex(pokemon string, config *PokeAPIConfig, pokedex *map[string]Pokemon) error {
	if len(*pokedex) == 0 {
		fmt.Println("You have not caught any pokemon yet!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range *pokedex {
		fmt.Println(" - " + key)
	}
	return nil
}
