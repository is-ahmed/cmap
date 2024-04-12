# Command Map

Command Map is a command-line interface (CLI) tool written in Go that helps you keep track of commonly used commands. With Command Map, you can easily store and retrieve frequently used commands, making it convenient to access them whenever you need them.

## Features

- **Store Commands**: Save frequently used commands along with a descriptive alias for easy reference.
- **Retrieve Commands**: Quickly retrieve stored commands using their aliases.
- **List Commands**: View a list of all stored commands along with their aliases.
- **Clear Commands**: Remove all stored commands to start fresh.

## Installation

To install Command Map, you can download the precompiled binary from the [Releases](https://github.com/yourusername/command-map/releases) page or build it from source.

```bash
# Install from source
go get -u github.com/is-ahmed/cmap

```

## Usage

- Initialize the mapping

```
> cmap init
```

- Add a command

```
> cmap insert
Pass a command: rm <file>
Pass a description: Remove file from the current working directory

```

- Search through existing commands

```
> cmap search
```

- List all commands

```
> cmap list
```

- Clear all commands from the map

```
> cmap clear
```


