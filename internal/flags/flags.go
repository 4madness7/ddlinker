package flags

import (
	"fmt"
	"os"
)

type Flag struct {
	description string
	shortName   string
	longName    string
	marked      bool
	Value       bool
}

func NewFlag(shortName rune, longName, description string, value bool) Flag {
	return Flag{
		description: description,
		longName:    "--" + longName,
		shortName:   "-" + string(shortName),
		marked:      false,
		Value:       value,
	}
}

type flags map[string]*Flag

var allFlags flags = make(flags)

func Register(f *Flag) error {
	_, exists := allFlags[f.shortName]
	if exists {
		return fmt.Errorf("Flag with short name of '%s' already exists.", string(f.shortName))
	}
	_, exists = allFlags[f.longName]
	if exists {
		return fmt.Errorf("Flag with long name of '%s' already exists.", f.longName)
	}
	allFlags[f.shortName] = f
	allFlags[f.longName] = f
	return nil
}

func Parse() ([]string, error) {
	args := os.Args[1:]
	i := 0
	for {
		if i >= len(args) {
			break
		}
		if args[i][0] != 45 {
			break
		}
		flag, exists := allFlags[args[i]]
		if !exists {
			return args, fmt.Errorf("Flag '%s' does not exist.", args[i])
		}
		if flag.marked {
			return args, fmt.Errorf("Duplicate flags '%s' and '%s'. Use only one.", flag.shortName, flag.longName)
		}
        flag.Value = true
        flag.marked = true
        i++
	}
	return args[i:], nil

}
