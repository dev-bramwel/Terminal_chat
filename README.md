```markdown
# Net-Cat: TCP Chat Server

Net-Cat is a TCP chat server written in Go. It allows multiple users to connect through Netcat, choose a name, send messages, receive broadcasts from other connected users, and see previous chat history when joining.

## Authors

- Moses Amani
- Margaret Apiyo
- Bramwel Mutugi

## Project Structure

```text
.
├── cmd/
│   └── server/
│       ├── main.go            # Application entry point
│       └── main_test.go       # Main package tests
├── internal/
│   └── chat/
│       ├── ascii.go           # Welcome banner
│       ├── ascii_test.go      # Welcome banner tests
│       ├── client.go          # Client connection and input handling
│       ├── message.go         # Message model and formatting helpers
│       ├── message_test.go    # Message formatting tests
│       ├── server.go          # TCP server, client list, history, and broadcasts
│       └── server_test.go     # Server behavior tests
├── pkg/
│   └── config/
│       ├── config.go          # Command-line configuration parsing
│       └── config_test.go     # Configuration tests
├── go.mod                     # Go module definition
├── Makefile                   # Build, run, test, and cleanup commands
├── insights.md                # Project notes
└── README.md                  # Project documentation

```

## Requirements

* Go 1.20 or newer
* Netcat, available as `nc`

## Run A Server

Run the server on the default port, `8989`:

```bash
make run

```

Run the server on a custom port:

```bash
make run-port PORT=2525

```

You can also run the server directly with Go:

```bash
go run ./cmd/server
go run ./cmd/server 2525

```

When the server starts successfully, it listens for incoming TCP clients on the selected port.

## Join A Server

### From the Same Computer (Localhost)

Open another terminal and connect with Netcat.

Join the default server:

```bash
nc localhost 8989

```

Join a server running on a custom port:

```bash
nc localhost 2525

```

### From a Different Computer (Remote Connection)

To connect to a server running on a different computer within the same network:

1. Find the local IP address of the server machine (e.g., using `ipconfig` on Windows or `ip a` / `ifconfig` on Linux/macOS).
2. Use Netcat on the client computer, replacing `localhost` with the server's actual IP address:

```bash
nc 192.168.1.50 8989

```

> **Note:** Ensure that the server machine's firewall allows incoming TCP traffic on the designated port (e.g., `8989` or `2525`).

After connecting, enter your name when prompted:

```text
[ENTER YOUR NAME]: Alice

```

Type a message and press Enter to send it to the other connected clients.

## Chat Commands

### Changing Your Name

You can change your chat name at any time while connected by using the `/name` command:

```text
/name NewName

```

This will update your identity in the chat and broadcast a system notification to all other connected clients.

## Exit A Server

To leave the chat from a client terminal, press:

```text
Ctrl+C

```

You can also close the client input stream with:

```text
Ctrl+D

```

To stop the server itself, go to the terminal running the server and press:

```text
Ctrl+C

```

## Concurrency Mechanics (Goroutines, Channels, and Mutexes)

This project relies heavily on Go's primitive concurrency tools to handle multiple incoming client connections seamlessly without blocking execution.

### Goroutines

Goroutines are used to handle separate execution paths for listening and interacting with clients asynchronously:

* **Server Listening:** `go s.acceptConnections(listener)` runs in the background to constantly monitor and accept new TCP handshakes.
* **Client Handling:** `go client.Handle()` handles the unique scanner loops, command executions, and incoming messages for every single connected user simultaneously.

### Channels

Channels serve as a thread-safe pipeline to transfer socket data safely between independent goroutines:

* **Connection Pipeline:** `joinChan chan net.Conn` is utilized to pass newly accepted network connections from the background listening loop directly into the main server lifecycle for client initialization.

### Mutexes

Because multiple client goroutines continuously access and modify shared server data simultaneously, a `sync.Mutex` (`s.mu`) is integrated to avoid critical race conditions:

* **State Protection:** Safe lock boundaries prevent concurrent write collisions when adding/removing users from the active clients map, appending elements to the global `s.history` chat slice, or reading the client state map to push a broadcast.

## Useful Commands

Build the binary:

```bash
make build

```

Run all tests:

```bash
make test

```

Run tests with the race detector:

```bash
make race

```

Format Go files:

```bash
make fmt

```

Run vet checks:

```bash
make vet

```

Remove build output:

```bash
make clean

```

Show all available Makefile commands:

```bash
make help

```

## Features

* Accepts TCP clients through Netcat
* Supports up to 10 connected clients
* Prompts each client for a name when they join
* Broadcasts chat messages to other connected clients
* Sends previous chat history to new clients
* Announces when clients join or leave
* Protects shared server state with a mutex

```

```
