# azsub

Interactive Azure subscription switcher with fzf integration.

## Installation

```bash
go install github.com/ivanovaleksandar/azsub@latest
```

## Setup

### Linux & macOS

Add to `~/.bashrc` or `~/.zshrc`:
```bash
azsub() {
    eval $(command azsub "$@")
}
```

Or for Fish shell, add to `~/.config/fish/config.fish`:
```fish
function azsub
    eval (command azsub $argv)
end
```

### Windows (PowerShell)

Add to your PowerShell profile:
```powershell
function azsub {
    Invoke-Expression (& azsub.exe $args)
}
```

## Usage

```bash
azsub
```

Selects subscription with fzf (if installed) and sets `ARM_SUBSCRIPTION_ID` and `ARM_SUBSCRIPTION_NAME` environment variables.

## Requirements

- Azure CLI (`az`)
- fzf (optional, falls back to list mode)
