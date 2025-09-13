package main

import (
	pokeapi "github.com/amarquezmazzeo/pokego/internal/pokeapi"
	pokecache "github.com/amarquezmazzeo/pokego/internal/pokecache"
)

type configCommand struct {
	nextURL     *string
	previousURL *string
	cache       *pokecache.Cache
	pokedex     map[string]pokeapi.Pokemon
}
