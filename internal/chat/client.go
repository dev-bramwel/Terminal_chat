// Package chat manages individual chat clients that connect to the chat via the `nc` tool.

package chat

// "bufio" - For reading lines of text efficiently from connections.
// "errors" - For creating simple error values without formatting.
// "net" - For network connections (the "pipe" between server and client).
// "strings" - For trimming whitespace from names and messages.
import (
	"bufio"  // Buffered I/O - reads/writes text efficiently
	"errors" // Simple error creation (errors.New)
	"net"    // Network connections (TCP sockets)
	"strings" // String tools (TrimSpace removes spaces from beginning/end)
)

type Client struct {
	conn    net.Conn        //net.Conn is an interface
	name    string          // The client's chosen name in the chat
	server  *Server        // Pointer to the server running this chat
	ch      chan string     // Their personal "inbox" - messages come here
	scanner *bufio.Scanner // Reads what they type, one line at a time
	writer  *bufio.Writer   // Writes messages back to their screen
}

//func  NewClient  will create a new Client and prepares them for chatting.
// This is called every time a new person connects to our server.
//
