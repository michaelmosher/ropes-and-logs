package main

import (
	"flag"
	"fmt"

	statemachine "github.com/michaelmosher/ropes-and-logs/pkg/state-machine"
)

var config struct {
	key   string
	value int
}

func init() {
	flag.StringVar(&config.key, "key", "", "some key")
	flag.IntVar(&config.value, "value", 0, "some value")
	flag.Parse()
}

func main() {
	fmt.Println("Hello, world!")
	store := statemachine.New()

	store.Apply(statemachine.SetCommand{Key: config.key, Value: config.value})
	fmt.Printf("Data constructed: %+v\n", store)
}
