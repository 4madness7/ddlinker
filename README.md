# Declarative Dotfiles Linker
I made this for myself, but since it might be useful to others, here's a bit of documentation.
This is a little tool to manage dotfiles in a declarative way.
`ddlinker` will parse the `.ddlinker_config.toml` inside your current working directory and create symlinks for you.

---
# Installation
To install you'll need [go](https://go.dev/doc/install) installed and run
`go install github.com/4madness7/ddlinker@latest`

---
# Usage
`ddlinker <flags> <command>`

| Flags        | Description                                                       | Usage                          |
| ------------ | ----------------------------------------------------------------- | ------------------------------ |
| -h,--help    | When used, prints help menu or description for specified command. | `ddlinker --help <command>`    |
| -v,--verbose | When used, ddlinker will print more detailed output.              | `ddlinker --verbose <command>` |

| Commands | Description                                                                                | Usage                      |
| -------- | ------------------------------------------------------------------------------------------ | -------------------------- |
| generate | Generates '.ddlinker_config.toml' file in current directory.                               | `ddlinker generate`        |
| link     | Creates symlinks based on the configuration provided. Use it with '-v' for absolute paths. | `ddlinker <flags> link`    |
| preview  | Shows a preview of the final links. Use it with '-v' for absolute paths.                   | `ddlinker <flags> preview` |

---
# Config file
`ddlinker` uses a `.ddlinker_config.toml` to configure all links.

Here's an example of what the file looks like:
```toml
[[destinations]] # this will add a new item to the destinations list
name = "config" # unique identifier for destination
path = "~/.config" # unique path for destination
links = [ # all the links in the current working directory to link
    "hypr",
    "nvim",
]
```

The config above will create the links as following:
```
path-to-dofiles/hypr -> ~/.config/hypr
path-to-dofiles/nvim -> ~/.config/nvim
```

The reason for this choice is because I like keeping my dotfiles like this
```
.
├── .zshrc # file
├── bin # dir
├── hypr # dir
└── nvim # dir
```
and with a tool like [stow](https://www.gnu.org/software/stow/manual/stow.html) I (think) couldn't really achieve that.

For the dir tree above, a config file might look something like this
```toml
[[destinations]] 
name = "config" 
path = "~/.config"
links = [ "hypr", "nvim" ]

[[destinations]] 
name = "home" 
path = "~"
links = [ ".zshrc" ]

[[destinations]] 
name = "local-bin" 
path = "~/.local"
links = [ "bin" ]
```
and `ddlinker` will create the following links
```
./.zshrc -> ~/.zshrc

./hypr -> ~/.config/hypr
./nvim -> ~/.config/nvim

./bin -> ~/.local/bin
```