# VGO-JUMP

`vgo-jump` is a simple CLI tool that helps locate the full namespace of a Laravel package or class by extracting the word at a specific position in a file and matching it against `use` statements.

## Features

- Extracts a word from a specified line and column in a file
- Finds the corresponding `use` statement to resolve the full package path
- Outputs the matched path to stdout
- Designed for Laravel PHP projects

## Usage

```bash
./vgo-jump <file> <line> <col>
```

## Build

```bash
go build -o vgo-jump main.go utils.go filepath_utils.go
```

## To use in VIM

```bash
cp vgo-jump /usr/local/bin/vgo-jump
```

In VIM:
```vim
:!vgo-jump %:p <line> <col>
```
Use this command in VIM to run `vgo-jump` on the current file at the specified line and column.

## Limitations

- Only works with basic `use` statements
- Assumes valid UTF-8 input files
- Assumes you have Alacritty installed
- Opens a new terminal instead of a new tab in VIM

## Future modifications

- [ ] Multiple languages support
- [ ] Open a new tab instead of a terminal
- [ ] Find a better way to get the project root path