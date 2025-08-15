# MCP Terminal Server

A simple MCP (Model Context Protocol) server implementation that provides a terminal command execution tool. This server allows clients to execute terminal commands remotely through the MCP protocol.

## Features

- Uses the official MCP Go SDK (`github.com/modelcontextprotocol/go-sdk`)
- Terminal tool for executing shell commands
- Stdio-based MCP server implementation (using standard input/output)
- Simple and lightweight design
- Graceful shutdown handling

## Getting Started

### Prerequisites

- Go 1.24 or higher (for local development)
- Docker (for containerized deployment)

### Local Installation

1. Clone the repository
2. Navigate to the project directory
3. Run the server:

```bash
go run cmd/server/main.go
```

The server communicates through standard input and output streams, making it suitable for integration with other applications that can interact via stdio.

### Docker Deployment

#### Building the Docker Image

You can build the Docker image using the provided build script:

```bash
./build.sh
```

Or manually with Docker:

```bash
docker build -t mcp-terminal-server:latest .
```

#### Running the Docker Container

Since this server communicates via standard input/output, you need to run it in interactive mode:

```bash
docker run -i --rm mcp-terminal-server:latest
```

You can pipe JSON requests to the container:

```bash
cat request.json | docker run -i --rm mcp-terminal-server:latest
```

#### Using Docker Compose

A `docker-compose.yml` file is provided for convenience. However, due to the nature of this service (communicating via stdin/stdout), it's typically not started with `docker-compose up`. Instead, you would build the image with docker-compose and then run it interactively:

```bash
# Build the image using docker-compose
docker-compose build

# Run the container interactively
docker run -i --rm mcp-terminal-server:latest
```

## MCP Protocol

This server implements the Model Context Protocol (MCP), which is a standardized protocol for communication between AI models and tools. The server exposes a terminal tool that can be used by MCP clients to execute shell commands through standard input/output streams.

### Terminal Tool

The terminal tool accepts the following input:

```json
{
  "command": "echo \"Hello from MCP Terminal Tool!\""
}
```

And returns a response with the following structure:

```json
{
  "output": "Hello from MCP Terminal Tool!\n",
  "exit_code": 0,
  "error": "" // Only present if there was an error
}
```

## Project Structure

- `cmd/server/main.go`: Server entry point
- `internal/terminal/terminal.go`: Terminal tool implementation

## Security Considerations

This implementation allows execution of arbitrary shell commands, which can be a security risk. In a production environment, you should implement proper authentication, authorization, and command validation to prevent misuse.

## License

This project is licensed under the MIT License - see the LICENSE file for details.