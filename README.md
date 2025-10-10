# Kite CLI - Infrastructure Template Manager

A CLI tool for managing infrastructure templates

## Features

- **Simple & Fast**: Download infrastructure templates with a single command
- **Version Tracking**: Keep track of installed templates and their versions
- **Clean Architecture**: DDD-inspired structure with clear separation of concerns
- **Local Config**: Uses `kite.json` in your project directory

## Installation

```bash
go build -o kite ./cmd/kite.go
```

## Quick Start

```bash
# Initialize a new project
kite init

# List available templates
kite list

# Add templates
kite add aws-vpc
kite add aws-vpc k8s-monitoring docker-redis  # multiple at once

# Check for updates
kite diff aws-vpc
kite update

# Remove templates
kite remove aws-vpc
```

## Commands

### `kite init`
Initialize a new project with kite.json config

```bash
kite init
kite init --path ./infra --registry https://custom-registry.com
```

### `kite add <template-name> [template-name...]`
Download and install templates from the registry

```bash
kite add aws-vpc
kite add aws-vpc k8s-monitoring  # install multiple
```

### `kite list`
List all available templates from the registry

```bash
kite list
```

### `kite remove <template-name>`
Remove an installed template

```bash
kite remove aws-vpc
```

### `kite update`
Check for and install updates for installed templates

```bash
kite update
```

### `kite diff <template-name>`
Show differences between local and registry versions

```bash
kite diff aws-vpc
```

## Configuration

The `kite.json` file is created in your project directory:

```json
{
  "version": "1.0.0",
  "registry": "https://api.kite.sh",
  "path": "./infrastructure",
  "templates": {
    "aws-vpc": {
      "version": "1.0.0",
      "installed": "2024-01-15T10:30:00Z"
    }
  }
}
```

## Project Structure

```
/cmd
  /kite              # Main entry point
/internal
  /cli               # CLI commands (init, add, list, remove, update, diff)
  /config            # Configuration management (kite.json)
  /container         # Dependency injection container
  /registry          # Registry client and types (API contract)
  /template          # Template installation and management logic
```

### Architecture Principles

- **Domain-Driven Design**: Clear domain boundaries (config, registry, template)
- **Single Responsibility**: Each file has one clear purpose (<100 lines)
- **Dependency Injection**: DI container manages all dependencies
- **Clean Separation**: CLI layer is thin, business logic in domain packages
- **Interface-based**: Easy to swap implementations (mock/real registry)

### Dependency Injection

The project uses a simple DI container pattern:

```go
// Container wires up all dependencies
container := container.New()

// Access managed dependencies
config := container.Config()
client := container.Client()
manager := container.Manager()
```

**Benefits:**
- Single place to manage dependency creation
- Easy to swap implementations (mock for testing)
- No circular dependencies
- Clear dependency graph

## Registry API Contract

### GET /templates
Returns list of available templates

```json
{
  "templates": [
    {
      "name": "aws-vpc",
      "description": "Production-ready AWS VPC",
      "version": "1.0.0",
      "tags": ["aws", "networking"],
      "author": "Kite Team"
    }
  ]
}
```

### GET /templates/{name}
Returns template details with files

```json
{
  "name": "aws-vpc",
  "version": "1.0.0",
  "description": "Production-ready AWS VPC",
  "files": [
    {
      "path": "main.tf",
      "content": "# VPC Configuration..."
    }
  ],
  "variables": [],
  "readme": "# AWS VPC Template..."
}
```

## Development

```bash
# Build
go build -o kite ./cmd/kite.go

# Run tests (when added)
go test ./...

# Install locally
go install ./cmd/kite.go
```

## Roadmap

- [x] Core commands (init, add, list, remove, update, diff)
- [x] Mock registry for testing
- [x] Local kite.json configuration
- [x] Colored output
- [ ] Real HTTP registry support
- [ ] File-based registry support
- [ ] Template variables/prompts
- [ ] Offline mode with cache
- [ ] Custom registry support
- [ ] Template search
- [ ] Dependency resolution

## License

MIT
