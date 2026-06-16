// Package chat manages individual chat clients that connect to the chat via the `nc` tool.

package chat

// "bufio" - For reading lines of text efficiently from connections.
// "errors" - For creating simple error values without formatting.
// "net" - For network connections (the "pipe" between server and client).
// "strings" - For trimming whitespace from names and messages.
import (
	"bufio"   // Buffered I/O - reads/writes text efficiently
	"errors"  // Simple error creation (errors.New)
	"fmt"     // Formats command responses
	"net"     // Network connections (TCP sockets)
	"strings" // String tools (TrimSpace removes spaces from beginning/end)
)

type Client struct {
	conn    net.Conn       //net.Conn is an interface
	name    string         // The client's chosen name in the chat
	server  *Server        // Pointer to the server running this chat
	ch      chan string    // Their personal "inbox" - messages come here
	scanner *bufio.Scanner // Reads what they type, one line at a time
	writer  *bufio.Writer  // Writes messages back to their screen
}

// func  NewClient  will create a new Client and prepares them for chatting.
// This is called every time a new person connects to our server.
func NewClient(conn net.Conn, server *Server) *Client {

	return &Client{
		conn:    conn,                   // Save the connection
		server:  server,                 // Save the server reference
		ch:      make(chan string, 100), // Create their message inbox
		scanner: bufio.NewScanner(conn), // Set up a scanner to read their messages
		writer:  bufio.NewWriter(conn),  // Set up a writer to send messages to them
	}
}

func (c *Client) Handle() {
	defer c.conn.Close()

	c.Write(GetLinuxLogo())
	c.Write("[ENTER YOUR NAME]: ")

	if !c.scanner.Scan() {
		return
	}

	name := strings.TrimSpace(c.scanner.Text())
	if err := c.SetName(name); err != nil {
		c.WriteLine(err.Error())
		return
	}

	if !c.server.AddClient(c) {
		c.WriteLine("Server is full.")
		return
	}

	c.server.SendHistory(c)
	c.server.Broadcast(NewSystemMessage(c.Name()+" has joined our chat."), c)
	c.Write(Prompt(c.Name()))
	c.Read()
}

func (c *Client) Name() string {
	return c.name
}

func (c *Client) SetName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("empty name")
	}

	c.name = name
	return nil
}

func (c *Client) Write(message string) {
	_, _ = c.writer.WriteString(message)
	_ = c.writer.Flush()
}

func (c *Client) WriteLine(message string) {
	c.Write(message + "\n")
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
		message := NewChatMessage(c.Name(), msg)
		c.server.SaveMessage(message)
		c.server.Broadcast(message, c)
		c.Write(Prompt(c.Name()))
	}

	// If we reach this line, the client disconnected (closed their terminal).
	// We tell the server to clean up and notify others.
	c.server.RemoveClient(c)
}

func (c *Client) handleCommand(command string) {
	fields := strings.Fields(command)
	if len(fields) == 0 {
		return
	}

	switch fields[0] {
	case "/name":
		if len(fields) < 2 {
			c.WriteLine("usage: /name <new-name>")
			c.Write(Prompt(c.Name()))
			return
		}

		oldName := c.Name()
		newName := strings.Join(fields[1:], " ")
		if err := c.SetName(newName); err != nil {
			c.WriteLine(err.Error())
			c.Write(Prompt(oldName))
			return
		}

		c.server.Broadcast(NewSystemMessage(fmt.Sprintf("%s is now known as %s.", oldName, c.Name())), c)
	default:
		c.WriteLine("unknown command")
	}

	c.Write(Prompt(c.Name()))
}
