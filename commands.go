package main

import (
	"fmt"
	"strings"
)

type Commands map[string]Command

type Command struct {
	name  string
	desc  string
	usage string
	run   func(*Data) error
}

func getCommands() Commands {
	return Commands{
		"preview": Command{
			name:  "preview",
			desc:  "Shows a preview of the final links.\nUse it with '-v' for absolute paths.",
			usage: "ddlinker <flags> preview",
			run:   previewHandler,
		},

		"generate": Command{
			name:  "generate",
			desc:  "Generates '.ddlinker_config.toml' file in current directory.",
			usage: "ddlinker generate",
			run:   generateHandler,
		},

		"link": Command{
			name:  "link",
			desc:  "Creates symlinks based on the configuration provided.\nUse it with '-v' for absolute paths.",
			usage: "ddlinker <flags> link",
			run:   linkHandler,
		},
	}
}

func (c Command) getHelp() string {
	lines := strings.Split(c.desc, "\n")
	desc := lines[0] + fmt.Sprintf("    Eg. %s\n", c.usage)
	for _, line := range lines[1:] {
		desc += fmt.Sprintf("\t\t%s\n", line)
	}
	return fmt.Sprintf("  %s%3s%s", c.name, "\t", desc)
}

func (c Commands) getOrderedKeys() []string {
	keys := make([]string, len(c))
	i := 0
	for k := range c {
		keys[i] = k
		i++
	}

	for i := 0; i < len(keys)-1; i++ {
		for j := i; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	return keys
}
