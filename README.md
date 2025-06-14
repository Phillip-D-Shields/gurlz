# Gurlz 🌐

> **Golang URLs** - A fast CLI tool for managing and testing HTTP requests

Gurlz (pronounced "girls") is a lightweight command-line tool that lets you save, organize, and execute HTTP requests. Think of it as a simplified Postman for your terminal.

## Features

- 🚀 **Fast & Lightweight**: Built with Go for speed
- 💾 **Persistent Storage**: Save requests locally in YAML format
- 🎯 **Simple Commands**: Intuitive CLI interface using Cobra
- 🔧 **Flexible Headers**: Support for custom headers and request bodies
- 📁 **Collections**: Organize requests into collections (coming soon)
- 🔄 **Hot Reload**: Development with Air for instant testing

## Installation

### From Source
```bash
git clone https://github.com/yourusername/gurlz.git
cd gurlz
go install .
```

### Development Setup
```bash
# Install Air for hot reloading
go install github.com/cosmtrek/air@latest

# Run with hot reload
air
```

## Quick Start

```bash
# Add a simple GET request
gurlz add api-health https://api.example.com/health

# Add a POST request with headers and body
gurlz add create-user https://api.example.com/users \
  -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer your-token" \
  -d '{"name":"John","email":"john@example.com"}'

# List all saved requests (coming soon)
gurlz list

# Execute a request (coming soon)
gurlz ping api-health
```

## Usage

### Adding Requests

```bash
# Basic syntax
gurlz add <n> <url> [flags]

# Examples
gurlz add github-api https://api.github.com/user
gurlz add post-data https://httpbin.org/post -X POST -d '{"key":"value"}'
gurlz add with-auth https://api.example.com/protected -H "Authorization: Bearer token123"
```

### Flags

- `-X, --method`: HTTP method (default: GET)
- `-H, --header`: HTTP headers (can be used multiple times)
- `-d, --data`: Request body data

## Configuration

Gurlz stores configuration and requests in `~/.gurlz/`:

```
~/.gurlz/
├── config.yaml      # Global settings
├── requests.yaml    # Saved requests
└── collections.yaml # Request collections (coming soon)
```

### Default Configuration

```yaml
default_headers:
  User-Agent: "gurlz/1.0.0"
timeout: "30s"
follow_redirect: true
save_responses: true
output_format: "json"
color_output: true
default_method: "GET"
max_response_size: 1048576
```

## Commands

| Command | Status | Description |
|---------|--------|-------------|
| `add` | ✅ | Add a new request |
| `list` | 🚧 | List all saved requests |
| `show` | 🚧 | Show request details |
| `ping` | 🚧 | Execute a request |
| `edit` | 🚧 | Edit an existing request |
| `remove` | 🚧 | Remove a request |
| `collection` | 🚧 | Manage collections |
| `config` | 🚧 | Manage configuration |

## Development

This project uses:
- **Cobra** for CLI framework
- **YAML** for configuration and data storage
- **UUID** for unique request identifiers
- **Air** for hot reloading during development

### Project Structure

```
gurlz/
├── main.go           # Entry point
├── cmd/              # Cobra commands
│   ├── root.go       # Root command
│   └── add.go        # Add command
├── internal/         # Internal packages
│   ├── models.go     # Data structures
│   └── storage.go    # File I/O operations
├── .air.toml         # Air configuration
├── .gitignore        # Git ignore rules
└── README.md         # This file
```

### Running Tests

```bash
go test ./...
```

### Building

```bash
# Build for current platform
go build -o gurlz .

# Build for multiple platforms
GOOS=linux GOARCH=amd64 go build -o gurlz-linux-amd64 .
GOOS=darwin GOARCH=amd64 go build -o gurlz-darwin-amd64 .
GOOS=windows GOARCH=amd64 go build -o gurlz-windows-amd64.exe .
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Roadmap

- [ ] `list` command - List all saved requests
- [ ] `ping` command - Execute requests and show responses
- [ ] `edit` command - Modify existing requests
- [ ] Collections support
- [ ] Response history
- [ ] Request monitoring
- [ ] Import/Export (Postman, curl)
- [ ] Configuration management
- [ ] Response assertions
- [ ] SQLite storage option

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) CLI framework
- Inspired by tools like curl, Postman, and HTTPie