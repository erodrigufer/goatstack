# goatstack

A project scaffolding tool for creating `Go + templ + htmx` web applications, similar to Create React App or Vite.

`goatstack` creates a full-stack web application in pure Go.
A single command generates all the required boilerplate to have a deployable project.
You can then start working right away in your application's business logic, on your views or HTTP handlers.

The generated Go backend uses exclusively the Go standard library, HTML templating is implemented with [templ](https://templ.guide/) and [htmx](https://htmx.org/) is vendored into the application.

## Prerequisites

Before using `goatstack`, ensure you have the following dependencies installed:

- Go 1.26.0+: Required for building the scaffolding tool and the generated projects
- just: Command runner

## Installation

```sh
just install
```

This builds the `goatstack` binary and copies it to `~/bin`.

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

After creation, run `just dev` in a directory within the generated project to start developing.

## Generated Project Structure

The generated Go web application has the following structure:

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
│   │   └── web/                # Web/HTTP utility functions
│   ├── go.mod                  # Go module file
│   └── .air.toml               # air live reload configuration
├── freebsd/
│   ├── INSTALL.sh              # FreeBSD installation script
│   └── myprojectd              # FreeBSD rc script
├── justfile
├── .envrc                      # Environment variables setup
├── .gitignore
└── README.md
```

### Backend

The HTTP server, handlers and middlewares are implemented using primarily the Go standard library.
It implements middleware functions to handle logging, security headers, authentication, session management and more.

`goatstack` can setup your project to work with either SQLite or PostgreSQL.

The package `daemonize` implements functions to make managing goroutines easier.

`emaild` provides a daemon to periodically send emails through SMTP.

## Developing goatstack

Use the following just targets for working on `goatstack`:

| Command        | Description                                           |
| -------------- | ----------------------------------------------------- |
| `just`         | List available targets                                |
| `just build`   | Build the goatstack binary                            |
| `just test`    | Run tests                                             |
| `just vet`     | Run go vet                                            |
| `just clean`   | Remove build and tmp folders                          |
| `just install` | Build and install to ~/bin                            |
| `just binary`  | Test the compiled binary by creating a sample project |

## Limitations

- The project currently only provides rc scripts and shell scripts to deploy to FreeBSD.
- The provided authentication middleware requires further customization by the user.
