# goatstack

Golang + Templ - Greatest Of All Time Stack - GOlAngTemplStack

A project scaffolding tool for creating Go + Templ web applications, similar to Create React App or Vite.

## Features

The created Go web applications have the following features:

- **Full-stack Go Web Applications**: Backend + Frontend in pure Go
  - **Templ Integration**: Modern HTML templating with Templ
  - **HTMX Integration**: Dynamic frontend without complex JavaScript
- **Database Support**: Integrated SQLite or PostgreSQL
- **Production Ready**: Includes daemonization and deployment scripts for FreeBSD
- **Development Tools**: Live reload with Air, comprehensive Justfile

## Prerequisites

Before using `goatstack`, ensure you have the following dependencies installed:

- **Go 1.26.0+**: Required for building the scaffolding tool and the generated projects
- **just**: Command runner (used for build automation)

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

| Flag       | Alias | Description                      | Required | Default  |
| ---------- | ----- | -------------------------------- | -------- | -------- |
| `--app`    | `-a`  | Main application name            | Yes      | -        |
| `--module` | `-m`  | Go module name for go.mod        | Yes      | -        |
| `--daemon` | `-d`  | Daemon name for deployment       | Yes      | -        |
| `--db`     | -     | Database type (sqlite, postgres) | No       | postgres |

### Example

```sh
goatstack create --app myproject --module codeberg.org/myproject --daemon myprojectd
```

After creation, run `just dev` in the generated project to start developing.

## Generated Project Structure

The scaffolding generates a comprehensive Go web application with the following structure:

```
myproject/
├── backend/
│   ├── cmd/
│   │   └── myproject/          # Main application
│   ├── internal/
│   │   ├── daemonize/          # Goroutine (daemon) management
│   │   ├── emaild/             # Email daemon
│   │   ├── server/             # HTTP server
│   │   ├── state/              # Database queries
│   │   ├── static/             # Static assets (CSS, JS)
│   │   ├── views/              # Templ components/views
│   │   └── web/                # Web utilities
│   ├── go.mod                  # Go module file
│   └── .air.toml               # Live reload config
├── freebsd/
│   ├── INSTALL.sh              # FreeBSD installation script
│   └── myprojectd              # FreeBSD rc script
├── justfile
├── .envrc                      # Environment setup
├── .gitignore
└── README.md
```

### Backend Features

- **HTTP Server** with middleware (logging, security headers, authentication, session management)
- **Database integration** (SQLite or PostgreSQL)
- **Goroutine management**
- **Email daemon** for sending emails with SMTP
- **Configuration management** via environment variables
- **Health checks** and basic endpoints

### Frontend Features

- **Templ components** for HTML templating
- **HTMX integration** for dynamic interactions withous JS
- **CSS styling** with base styles

### Development Tools

- **Air** for live reloading during development
- **Comprehensive Justfile** with targets for:
  - `just dev` - Start development server with live reloading
  - `just build` - Build production binary
  - `just test` - Run tests
  - `just deploy` - Deployment scripts

### Deployment

- **FreeBSD rc script** to create a system daemon
- **Installation script**

## Developing goatstack

Use the following just targets for working on `goatstack`:

| Command        | Description                                           |
| -------------- | ----------------------------------------------------- |
| `just`         | List available targets                                |
| `just build`   | Build the binary                                      |
| `just test`    | Run tests                                             |
| `just vet`     | Run go vet                                            |
| `just clean`   | Remove build and tmp folders                          |
| `just install` | Build and install to ~/bin                            |
| `just binary`  | Test the compiled binary by creating a sample project |

## Limitations

- Currently focused on FreeBSD deployment (rc scripts)
- Limited to SQLite and PostgreSQL databases
- Basic authentication middleware included but may need customization

## Future Plans (maybe?)

- Linux systemd service files?
- Docker support
- Additional Templ components and UI kits?
