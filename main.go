package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/Jschles1/pokedexcli/internal/pokeapi"
	"github.com/Jschles1/pokedexcli/internal/pokecache"
)

func main() {
	commands := getCommands()
	cache := pokecache.NewCache(5 * time.Minute)
	client := pokeapi.NewClient(cache)

	config := &PokeAPIConfig{
		Next:     "https://pokeapi.co/api/v2/location-area",
		Previous: "",
		Client:   client,
	}
	pokedex := make(map[string]Pokemon)
	for {
		fmt.Printf("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		cmd := strings.Fields(scanner.Text())
		commandName := cmd[0]
		commandParam := ""
		if len(cmd) > 1 {
			commandParam = cmd[1]
		}
		if command, ok := commands[commandName]; ok {
			command.callback(commandParam, config, &pokedex)
		} else {
			fmt.Println("Invalid command, please try again. (Type \"help\" for a list of valid commands)")
		}
	}
}
