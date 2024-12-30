package flags

import (
	"fmt"
	"os"
)

type Flag struct {
	description string
	shortName   string
	longName    string
	usage       string
	marked      bool
	Value       bool
}

func NewFlag(shortName rune, longName, description, usage string, value bool) *Flag {
	return &Flag{
		description: description,
		longName:    "--" + longName,
		shortName:   "-" + string(shortName),
		usage:       usage,
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

func GetHelpMenu() string {
	uniqueFlags := map[*Flag]struct{}{}
	for _, f := range allFlags {
		if _, ok := uniqueFlags[f]; !ok {
			uniqueFlags[f] = struct{}{}
		}
	}

	helpMenu := ""

	flags := make([]*Flag, len(uniqueFlags))
	i := 0
	for k := range uniqueFlags {
		flags[i] = k
		i++
	}
	for i := 0; i < len(flags)-1; i++ {
		for j := i; j < len(flags); j++ {
			if flags[i].shortName > flags[j].shortName {
				flags[i], flags[j] = flags[j], flags[i]
			}
		}
	}

	for _, v := range flags {
		helpMenu = helpMenu + fmt.Sprintf("  %s,%s\t%s  Eg. %s\n\n", v.shortName, v.longName, v.description, v.usage)
	}

	return helpMenu
}
