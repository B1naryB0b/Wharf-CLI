# Wharf-CLI

A cross-platform command-line tool for managing and inspecting network ports on Windows, Linux, and macOS.

## Overview

Wharf-CLI provides a unified interface for network port management across different operating systems. Instead of remembering platform-specific commands like `netstat`, `ss`, or `lsof`, use Wharf's simple commands to view port information consistently on any platform.

## Features

- **Cross-Platform Support**: Works seamlessly on Windows, Linux, and macOS
- **Simple Commands**: Easy-to-remember commands for common port operations
- **Port Inspection**: List all active ports or query specific port numbers
- **Process Information**: View which processes are using network ports
- **Port Checking**: Check connectivity to specific ports with configurable timeout
- **Free Port Discovery**: Find available ports for your applications

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

## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `list [port]` | `l`   | List all active ports or check a specific port |
| `check [port] [timeout]` | `c`   | Check port connectivity with optional timeout (seconds) |
| `free [count]` | `f`   | Find available ports (default: 1) |
| `version` | `v`   | Display version information |


## Usage Examples

```bash
# List all active network ports
wharf list
wharf l

# Check if port 3000 is in use
wharf list 3000

# Check if port 8080 is accepting connections
wharf check 8080
wharf t localhost:8080 5.0

# Find available ports
wharf free
wharf f 10

# Display version information
wharf version
wharf v
```

## Platform-Specific Behavior

Wharf abstracts platform differences through a unified interface. Port listing operations automatically use the appropriate system command:

| Platform | Command Used | Protocol Support |
|----------|-------------|------------------|
| **Windows** | `netstat -ano` | TCP, UDP |
| **Linux** | `ss -tulpn` | TCP, UDP, Unix sockets |
| **macOS** | `lsof -i -P -n` | TCP, UDP, IPv4, IPv6 |

Port checking (`check`) and free port discovery (`free`) use Go's native `net` package for cross-platform compatibility without requiring elevated privileges.

## Development

### Project Structure

```
Wharf-CLI/
├── main.go                        # Entry point
├── cmd/
│   ├── root.go                    # Root command definition
│   ├── list.go                    # List command implementation
│   ├── check.go                    # Check command implementation
│   └── free.go                    # Free port finder implementation
├── internal/
│   └── ports/
│       ├── manager.go             # Platform manager interface
│       ├── manager_common.go      # Cross-platform implementations
│       ├── manager_windows.go     # Windows-specific implementations
│       ├── manager_linux.go       # Linux-specific implementations
│       ├── manager_mac.go         # macOS-specific implementations
│       └── helper.go              # Base command execution helper
├── go.mod                         # Go module definition
├── go.sum                         # Dependency checksums
├── LICENSE                        # GNU GPL v3
└── README.md                      # This file
```

### Built With

- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go

### Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Roadmap

### Planned Features

**Port Management:**
- Port killing/termination (`ClosePort`)
- Firewall rule management (`OpenFirewall`, `CloseFirewall`)
- Port opening/binding (`OpenPort`)

**Monitoring & Logging:**
- Port activity logging (`StartPortLog`, `EndPortLog`, `GetActiveLogs`)
- Real-time port monitoring

**Portability:**
- Portable mode: Single binary with runtime platform detection for use on external storage devices (USB drives, etc.)

**Output & Configuration:**
- JSON/CSV output formats
- Configuration file support
- Port group management

## Support

For issues, questions, or contributions, please visit the [GitHub repository](https://github.com/B1naryB0b/Wharf-CLI).
