# goatstack

Golang + Templ - Greatest Of All Time Stack - GOlAngTemplStack

A project scaffolding tool for creating Go + Templ web applications, similar to Create React App or Vite.

## Installation

```sh
just install
```

This builds the binary and copies it to `~/bin`.

`~/bin` should be in your `PATH` for `goatstack` to be accessible from your shell.

## Usage

```sh
goatstack create --app <name> --module <module> --daemon <daemon> [--db <type>]
```

### Flags

| Flag | Alias | Description | Required | Default |
|------|-------|-------------|----------|---------|
| `--app` | `-a` | Main application name | Yes | - |
| `--module` | `-m` | Go module name for go.mod | Yes | - |
| `--daemon` | `-d` | Daemon name for deployment | Yes | - |
| `--db` | - | Database type (sqlite, postgres) | No | postgres |

### Example

```sh
goatstack create --app myproject --module codeberg.org/myproject --daemon myprojectd
```

After creation, run `just dev` in the generated project to start developing.

## Development

### Available just commands

| Command | Description |
|---------|-------------|
| `just` | List available targets |
| `just build` | Build the binary |
| `just test` | Run tests |
| `just vet` | Run go vet |
| `just clean` | Remove build and tmp folders |
| `just install` | Build and install to ~/bin |
| `just binary` | Test the compiled binary by creating a sample project |
