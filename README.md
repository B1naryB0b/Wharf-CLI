# Wharf-CLI

A cross-platform command-line tool for managing and inspecting network ports on Windows, Linux, and macOS.

## Overview

Wharf-CLI provides a unified interface for network port management across different operating systems. Instead of remembering platform-specific commands like `netstat`, `ss`, or `lsof`, use Wharf's simple commands to view port information consistently on any platform.

## Features

- **Cross-Platform Support**: Works seamlessly on Windows, Linux, and macOS
- **Simple Commands**: Easy-to-remember commands for common port operations
- **Port Inspection**: List all active ports or query specific port numbers
- **Process Information**: View which processes are using network ports

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/B1naryB0b/Wharf-CLI.git
cd Wharf-CLI
```

2. Build the binary:
```bash
go build -o wharf
```

3. (Optional) Add to your PATH for global access

### Prerequisites

- Go 1.25.3 or higher

## Usage

### List All Ports

View all active network ports on your system:

```bash
wharf list
```

Alias:
```bash
wharf l
```

### Check Specific Port

Check if a specific port is in use:

```bash
wharf list 8080
```

Or using the alias:
```bash
wharf l 3000
```

## Platform-Specific Behavior

Wharf automatically uses the appropriate system command based on your platform:

- **Windows**: Uses `netstat -ano`
- **Linux**: Uses `ss -tulpn`
- **macOS**: Uses `lsof -i -P -n`

## Commands

| Command | Alias | Description | Example |
|---------|-------|-------------|---------|
| `wharf list` | `wharf l` | List all active ports | `wharf list` |
| `wharf list [port]` | `wharf l [port]` | Check specific port | `wharf list 8080` |

## Examples

```bash
# List all active network ports
wharf list

# Check if port 3000 is in use
wharf list 3000

# Using the short alias
wharf l 8080
```

## Development

### Project Structure

```
Wharf-CLI/
├── main.go           # Entry point
├── cmd/
│   ├── root.go       # Root command definition
│   ├── list.go       # List command implementation
│   └── wharf.go      # Port listing logic
├── go.mod            # Go module definition
├── go.sum            # Dependency checksums
├── LICENSE           # GNU GPL v3
└── README.md         # This file
```

### Built With

- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go

### Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Roadmap

Future features under consideration:
- Port killing/termination
- Port forwarding management
- Port group management
- Configuration file support
- JSON/CSV output formats

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/B1naryB0b/Wharf-CLI).
