package main

import (
	"fmt"
	"os"

	"github.com/4madness7/ddlinker/internal/config"
	"github.com/4madness7/ddlinker/internal/flags"
)

type Data struct {
	cfg   *config.Config
	flags struct {
		verbose *flags.Flag
		help    *flags.Flag
	}
}

func main() {
	data := Data{}
	data.flags.verbose = flags.NewFlag('v', "verbose", "palle", false)
	err := flags.Register(data.flags.verbose)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data.flags.help = flags.NewFlag('h', "help", "palle", false)
	err = flags.Register(data.flags.help)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	args, err := flags.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println("===== DEBUG =====")
	// fmt.Printf("help: %v | verb: %v\n", data.flags.help.Value, data.flags.verbose.Value)
	// fmt.Println(args)
	// fmt.Println("===== DEBUG =====")

	if len(args) == 0 && !data.flags.help.Value {
		fmt.Println("Please provide a command.")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Println("Too many commands.")
		os.Exit(1)
	}
	if data.flags.verbose.Value && data.flags.help.Value {
		fmt.Println("'verbose' and 'help' flag cannot be used together.")
		os.Exit(1)
	}

	input := ""
	if len(args) > 0 {
		input = args[0]
	}

	commands := getCommands()
	if data.flags.help.Value {
		err := helpHandler(commands, input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	cmd, ok := commands[input]
	if !ok {
		fmt.Printf("Command '%s' does not exist.\n", input)
		os.Exit(1)
	}

	data.cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	allWarns, allErrs := data.cfg.Validate()
	if len(allWarns) > 0 {
		fmt.Println("===== WARNINGS =====")
		fmt.Println("The following warnings were found while validating the config file:")
		for name, warns := range allWarns {
			fmt.Printf("-- Destination '%s' --\n", name)
			for _, err := range warns {
				fmt.Println("    -", err)
			}
			fmt.Println()
		}
	}
	if len(allErrs) > 0 {
		fmt.Println("===== ERRORS =====")
		fmt.Println("The following errors were found while validating the config file:")
		for name, errs := range allErrs {
			fmt.Printf("-- Destination '%s' --\n", name)
			for _, err := range errs {
				fmt.Println("    -", err)
			}
			fmt.Println()
		}
		os.Exit(1)
	}

	err = cmd.run(&data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
