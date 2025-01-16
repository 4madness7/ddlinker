package main

const (
	configCreationString = `# To use ddlinker you must create a 'destination'.
# Copy the following template to declare a new destination.
#
# [[destinations]]
# name = "config"
# path = "~/.config"
# links = ["nvim"]
#
# Every destination is saved in the 'destinations' array, that's why there is a
# [[destinations]] before all the info about the destination.
# 'name' is the name of the destination, it MUST be unique and can be any string you prefer.
# 'path' is the path of the destination, it's the directory where all the symlinks
# will be made. It MUST be unique and can be any string you prefer.
# 'links' are all the file/dirs that you want to be symlinked to the specified 'path'.
# Every string in the list must be a valid file/dir in your current directory.
# Take the destination above, it will create a symlink that will look something
# like this './nvim -> ~/.config/nvim'.
# For more info: https://github.com/4madness7/ddlinker`

	shortHelpMsg = "Use 'ddlinker --help' for more information."
)
