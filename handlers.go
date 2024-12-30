package main

import (
	"fmt"
	"os"
	"os/exec"
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
		fmt.Println(command.getHelp())
		return nil
	}
	fmt.Println("Flags")
	fmt.Println(flags.GetHelpMenu())
	fmt.Println("Commands")
	for _, k := range commands.getOrderedKeys() {
		fmt.Println(commands[k].getHelp())
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

func linkHandler(data *Data) error {
	for _, dest := range data.cfg.Destinations {
		path, err := getFullPath(dest.Path)
		if err != nil {
			return err
		}
		title := fmt.Sprintf("Linking '%s' | Path: ", dest.Name)
		if data.flags.verbose.Value {
			title += fmt.Sprintf("'%s'", path)
		} else {
			title += fmt.Sprintf("'%s'", dest.Path)
		}
		fmt.Println(title)
		for _, link := range dest.Links {
			fullLink, err := filepath.Abs(link)
			if err != nil {
				return err
			}
			destFullPath := filepath.Join(path, link)
			var linkMsg string
			if data.flags.verbose.Value {
				linkMsg = fmt.Sprintf("%s -> %s", fullLink, destFullPath)
			} else {
				linkMsg = fmt.Sprintf("./%s -> %s", link, filepath.Join(dest.Path, link))
			}
			cmd := exec.Command("ln", "-s", fullLink, destFullPath)
			_, err = cmd.Output()
			if err != nil {
				fmt.Printf("  ERROR: File/Directory already exists | %s", linkMsg)
				continue
			}
			fmt.Printf("  SUCCESSFUL | %s", linkMsg)
		}
		fmt.Println()
		fmt.Println()
	}
	return nil
}
