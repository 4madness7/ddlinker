package main

import (
	"fmt"
	"os"

	"github.com/4madness7/ddlinker/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := range cfg.Destinations {
		fmt.Printf(
			"Name \t%s\nPath \t%s\nLinks \t%v\n\n",
			cfg.Destinations[i].Name,
			cfg.Destinations[i].Path,
			cfg.Destinations[i].Links,
		)
	}
}
