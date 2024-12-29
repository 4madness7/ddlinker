package main

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
			desc:  "Shows a preview of the final links. Use it with '-v' for full paths.",
			usage: "ddlinker preview",
			run:   previewHandler,
		},

		"generate": Command{
			name:  "generate",
			desc:  "Generates '.ddlinker_config.toml' file in current directory.",
			usage: "ddlinker generate",
			run:   generateHandler,
		},
	}
}
