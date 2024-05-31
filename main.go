package main

import "github.com/Muhamaddiis/pokedec-cli/internal/pokeapi"

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}
	startRepl(&cfg)

}
