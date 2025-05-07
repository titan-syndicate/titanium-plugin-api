# Titanium Plugin API SDK

This SDK provides the necessary tools and interfaces for developing plugins for the Titanium platform. It includes the RPC interfaces and a structured logging system.

## Installation

```bash
go get github.com/titan-syndicate/titanium-plugin-sdk
```

## Usage

### In Your Plugin

```go
package main

import (
    "github.com/titan-syndicate/titanium-plugin-api/pkg/logger"
    "github.com/titan-syndicate/titanium-plugin-api/pkg/pluginapi"
)

func main() {
    // Initialize logger
    if err := logger.Init(logger.Config{
        Level:      "info",
        PluginName: "my-plugin",
    }); err != nil {
        panic(err)
    }
    defer logger.Sync()

    // Your plugin implementation here
}
```

## Development

### Prerequisites

- Go 1.24 or later
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```
- Mage build tool:
  ```bash
  go install github.com/magefile/mage@latest
  ```

### Building

This project uses [Mage](https://magefile.org/) for build automation. Available commands:

```bash
# Generate protobuf code
mage generate

# Clean generated files
mage clean

# Clean and regenerate everything
mage all
```

### Project Structure

```
titanium-plugin-api/
├── pkg/
│   ├── pluginapi/     # RPC interfaces and protobuf definitions
│   └── logger/        # Structured logging package
├── magefile.go        # Build automation
└── README.md
```

### Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

## License

[Add your license here]
