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
func NewClient(conn net.Conn, server *Server) *Client {
	
	return &Client{
		conn:    conn,               // Save the connection
		server:  server,             // Save the server reference
		ch:      make(chan string, 100), // Create their message inbox
		scanner: bufio.NewScanner(conn),  // Set up a scanner to read their messages
		writer:  bufio.NewWriter(conn),   // Set up a writer to send messages to them
	}
}

func (c *Client) Read() {
	// for c.scanner.Scan() loops until the client disconnects.
	// scanner.Scan() reads one line of text and returns true if successful.
	// When the client closes their terminal or types Ctrl+C, it returns false.
	for c.scanner.Scan() {
		// .Text() gets the line they typed (as a string).
		msg := strings.TrimSpace(c.scanner.Text())

		
		if msg == "" {
			continue // Skip to the next message
		}

		// Check if this is a special command (like /name).
		// Commands start with "/" and do special actions instead of normal chatting.
		if strings.HasPrefix(msg, "/") {
			c.handleCommand(msg) // Process the command (like changing name)
			continue             // Don't broadcast commands as regular messages
		}

		// Broadcast the message to all other clients in the chat.
		// c.server.broadcast does the heavy lifting of sending it to everyone else.
		c.server.broadcast(msg, c)
	}

	// If we reach this line, the client disconnected (closed their terminal).
	// We tell the server to clean up and notify others.
	c.server.removeClient(c)
}
