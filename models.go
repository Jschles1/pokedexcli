package main

import (
	"github.com/Jschles1/pokedexcli/internal/pokeapi"
)

type PokeAPIConfig struct {
	Next     string
	Previous string
	Client   *pokeapi.Client
}

type Pokemon struct {
	name   string
	height int
	weight int
	stats  []PokemonStat
	types  []string
}

type PokemonStat struct {
	name  string
	value int
}
