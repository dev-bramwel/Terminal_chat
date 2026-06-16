# Net-Cat: TCP Chat Server

Net-Cat is a TCP chat server written in Go. It allows multiple users to connect through Netcat, choose a name, send messages, receive broadcasts from other connected users, and see previous chat history when joining.

## Authors

- [Moses Amani](https://learn.zone01kisumu.ke/git/mamani)
- [Margaret Apiyo](https://learn.zone01kisumu.ke/git/margaapiyo)
- [Bramwel Mutugi](https://learn.zone01kisumu.ke/git/mumutugi)

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

- Go 1.20 or newer
- Netcat, available as `nc`

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

Open another terminal and connect with Netcat.

Join the default server:

```bash
nc localhost 8989
```

Join a server running on a custom port:

```bash
nc localhost 2525
```

After connecting, enter your name when prompted:

```text
[ENTER YOUR NAME]: Alice
```

Type a message and press Enter to send it to the other connected clients.

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

- Accepts TCP clients through Netcat
- Supports up to 10 connected clients
- Prompts each client for a name when they join
- Broadcasts chat messages to other connected clients
- Sends previous chat history to new clients
- Announces when clients join or leave
- Protects shared server state with a mutex
