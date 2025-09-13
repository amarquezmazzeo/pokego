package main

import pokecache "github.com/amarquezmazzeo/pokego/internal/pokecache"

type configCommand struct {
	nextURL     *string
	previousURL *string
	cache       *pokecache.Cache
}
