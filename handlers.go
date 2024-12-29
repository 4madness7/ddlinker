package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/4madness7/ddlinker/internal/config"
	"github.com/4madness7/ddlinker/internal/flags"
)

func helpHandler(commands Commands, input string) error {
	fmt.Println("===== Help Menu =====")
	fmt.Println()
	if input != "" {
		command, ok := commands[input]
		if !ok {
			return fmt.Errorf("Command '%s' does not exist.", input)
		}
		fmt.Printf("%s\t%s Eg. '%s'\n", command.name, command.desc, command.usage)
		return nil
	}
	fmt.Println("Flags")
	fmt.Println(flags.GetHelpMenu())
	fmt.Println("Commands")
	for _, command := range commands {
		fmt.Printf("  %s\t%s Eg. '%s'\n", command.name, command.desc, command.usage)
	}
	return nil
}

func previewHandler(data *Data) error {
	for _, dest := range data.cfg.Destinations {
		var path string
		if data.flags.verbose.Value {
			var err error
			path, err = getFullPath(dest.Path)
			if err != nil {
				return err
			}
		} else {
			path = dest.Path
		}
		fmt.Printf("Destination name: %s\n", dest.Name)
		fmt.Printf("Destination path: %s\n", path)
		fmt.Println("Preview:")
		for _, link := range dest.Links {
			var fullLink string
			if data.flags.verbose.Value {
				var err error
				fullLink, err = filepath.Abs(link)
				if err != nil {
					return err
				}
			} else {
				fullLink = "./" + link
			}
			destFullPath := filepath.Join(path, link)
			fmt.Printf("  %s -> %s\n", fullLink, destFullPath)
		}
		fmt.Println()
	}
	return nil
}

func generateHandler(data *Data) error {
	fullPath, err := filepath.Abs(config.ConfigFileName)
	if err != nil {
		return err
	}
	_, err = os.Stat(fullPath)
	exists := !os.IsNotExist(err)
	if exists {
		return fmt.Errorf("Config file '%s' already exists.", config.ConfigFileName)
	}
	f, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	_, err = f.Write([]byte(configCreationString))
	if err != nil {
		return err
	}
	fmt.Printf("Config file '%s' created.\n", config.ConfigFileName)
	return nil
}
