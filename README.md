# Declarative Dotfiles Linker
This is `ddlinker` a cli tool to manage your dotfiles... Declaratively!
`ddlinker` will parse the `.ddlinker_config.toml` inside your current working
directory and create symlinks for you.

# Why
I really like the declarative nature of tools like [Nix and NixOS](https://nixos.org/),
but I don't like the way they manage dotfiles for my linux system(s).
So after a bit of time a decided to create `ddlinker` to manage my dotfiles the
way I like it and now it's my go tool to do that!

# Installation
To install, you'll need [go](https://go.dev/doc/install) installed and run
```bash
go install github.com/4madness7/ddlinker@latest
```

# Usage
```bash
ddlinker <flags> <command>
```

| Flags        | Description                                                       | Usage                          |
| ------------ | ----------------------------------------------------------------- | ------------------------------ |
| -h,--help    | When used, prints help menu or description for specified command. | `ddlinker --help <command>`    |
| -v,--verbose | When used, ddlinker will print more detailed output.              | `ddlinker --verbose <command>` |

| Commands | Description                                                                                | Usage                      |
| -------- | ------------------------------------------------------------------------------------------ | -------------------------- |
| generate | Generates '.ddlinker_config.toml' file in current directory.                               | `ddlinker generate`        |
| link     | Creates symlinks based on the configuration provided. Use it with '-v' for absolute paths. | `ddlinker <flags> link`    |
| preview  | Shows a preview of the final links. Use it with '-v' for absolute paths.                   | `ddlinker <flags> preview` |

# Examples
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

[[destinations]]
name = "home"
path = "~"
links = [ ".zshrc" ]

[[destinations]]
name = "local-bin"
path = "~/.local"
links = [ "bin" ]
```

The config above will create the links as following:
```
./.zshrc -> ~/.zshrc

./hypr -> ~/.config/hypr
./nvim -> ~/.config/nvim

./bin -> ~/.local/bin
```

The `preview` command will give an output like this:
```
Destination name: home
Destination path: ~
Preview:
  ./.zshrc -> ~/.zshrc

Destination name: config
Destination path: ~/.config
Preview:
  ./hypr -> ~/.config/hypr
  ./nvim -> ~/.config/nvim

Destination name: local-bin
Destination path: ~/.local
Preview:
  ./bin -> ~/.local/bin
```

The `link` command will give an output like this:
```
`Linking 'home' | Path: '~'
  Done | ./.zshrc -> ~/.zshrc # if linked correctly

Linking 'config' | Path: '~/.config'
  Already linked | ./hypr -> ~/.config/hypr # if link is already present
  Error: file/dir already exists | ./nvim -> ~/.config/nvim # if there is already a file/dir in the destination

Linking 'local-bin' | Path: '~/.local'
  Error: destination is a symlink to a different file/dir | ./bin -> ~/.local/bin # if destination is a symlink, but points to a different file/dir
```

## Contributing

To contribute, feel free to a new branch and send in a pull request.

All PRs should be submitted to the `main` branch.
