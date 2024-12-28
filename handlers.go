package main

import (
	"fmt"
	"path/filepath"
)

func helpHandler(commands Commands, input string) error {
	if input != "" {
		command, ok := commands[input]
		if !ok {
			return fmt.Errorf("Command '%s' does not exist.\n", input)
		}
		fmt.Printf("%s\t\t%s    Eg. '%s'\n", command.name, command.desc, command.usage)
	}
	for _, command := range commands {
		fmt.Printf("%s\t\t%s    Eg. '%s'\n", command.name, command.desc, command.usage)
	}
	return nil
}

func previewHandler(data *Data) error {
	for _, dest := range data.cfg.Destinations {
		fmt.Printf("Destination name: %s\n", dest.Name)
		fmt.Printf("Destination path: %s\n", dest.Path)
		fmt.Println("Preview:")
		for _, link := range dest.Links {
			destFullPath := filepath.Join(dest.Path, link)
			fmt.Printf("%s -> %s\n", link, destFullPath)
		}
		fmt.Println()
	}
	return nil
}
