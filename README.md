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

## Limitations

- Only works with basic use statements

- Assumes valid UTF-8 input files