# VGO-JUMP

`vgo-jump` is a simple CLI tool that helps locate the full namespace of a Laravel package or class by extracting the word at a specific position in a file and matching it against `use` statements.

## Features

- Extracts a word from a specified line and column in a file
- Finds the corresponding `use` statement to resolve the full package path
- Outputs the matched path to stdout
- Designed for PHP, Python, Java projects

## Usage

```bash
./vgo-jump <file> <line> <col>
```

## Build

```bash
go build -o vgo-jump main.go utils.go filepath_utils.go language_config.go
```

## To use in VIM

```bash
sudo cp vgo-jump /usr/local/bin/vgo-jump
sudo mkdir -p /etc/vgo-jump 
sudo cp languages.yaml /etc/vgo-jump/languages.yaml
```

```bash
mkdir -p ~/.config/nvim/lua/vgojump
sudo nvim ~/.config/nvim/lua/vgojump/init.lua
```

```lua
local M = {}

M.jump = function()
  local filepath = vim.fn.expand("%:p")
  local line = vim.fn.line(".")
  local col = vim.fn.col(".")

  local cmd = string.format("vgo-jump %s %d %d", filepath, line, col)

  vim.fn.jobstart(cmd, {
    stdout_buffered = true,
    stderr_buffered = true,
    on_stdout = function(_, data, _)
      if data and #data > 0 then
        for _, output in ipairs(data) do
          if output and output ~= "" then
            vim.cmd("tabnew " .. output)
            break
          end
        end
      end
    end,
    on_stderr = function(_, data, _)
      if data then
        vim.notify(table.concat(data, "\n"), vim.log.levels.ERROR)
      end
    end,
  })
end

return M
```

```bash
sudo nvim ~/.config/nvim/init.lua
```

```lua
local status, vgojump = pcall(require, "vgojump")
if not status then
  vim.notify("vgojump module error: " .. vgojump, vim.log.levels.ERROR)
  return
end

vim.api.nvim_create_user_command("Jump", function()
  vgojump.jump()
end, {})

-- Set key mapping
vim.keymap.set("n", "<leader>j", ":Jump<CR>", { noremap = true, silent = true })
```

In VIM:
```vim
:Jump
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