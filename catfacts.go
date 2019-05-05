package main

import (
	"fmt"
	"os"

	"github.com/amccarthy1/catfacts"
)

const (
	commandFact   = "fact"
	commandFacts  = "facts"
	commandBreeds = "breeds"
)

var commandMap = map[string]func(){
	commandFact:   fact,
	commandFacts:  facts,
	commandBreeds: breeds,
}

func fact() {
	c := catfacts.NewClient()
	fact, err := c.GetRandomFact()
	if err != nil {
		panic(err)
	}
	fmt.Println(fact.Fact)
}

func breeds() {
	c := catfacts.NewClient()
	breeds, err := c.ListAllBreeds()
	if err != nil {
		panic(err)
	}
	for _, breed := range breeds {
		fmt.Println(breed.Breed)
	}
}

func facts() {
	c := catfacts.NewClient()
	facts, err := c.ListAllFacts()
	if err != nil {
		panic(err)
	}
	for _, fact := range facts {
		fmt.Println(fact.Fact)
	}
}

func usage() {
	fmt.Println("Usage: catfacts [subcommand]\nwhere subcommand is one of:")
	for key := range commandMap {
		fmt.Printf("  %s\n", key)
	}
	os.Exit(127)
}

func main() {
	args := os.Args
	if len(args) == 1 {
		usage()
	}
	cmd, ok := commandMap[args[1]]
	if !ok {
		fmt.Printf("Invalid subcommand '%s'\n", args[1])
		usage()
	}
	cmd()
}
